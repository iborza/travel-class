// Package handlers contains the full set of handler functions and routes
// supported by the web api.
package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgraph-io/travel/business/data"
	"github.com/dgraph-io/travel/business/data/city"
	"github.com/dgraph-io/travel/business/data/schema"
	"github.com/dimfeld/httptreemux/v5"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger, gqlConfig data.GraphQLConfig) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	c := custom{
		log:       log,
		gqlConfig: gqlConfig,
	}
	mux.Handle(http.MethodPost, "/upload", c.upload)

	return mux
}

type custom struct {
	log       *log.Logger
	gqlConfig data.GraphQLConfig
}

func (c *custom) upload(w http.ResponseWriter, r *http.Request) {
	c.log.Println("upload: request:", r)

	var request schema.UploadFeedRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		c.log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	gql := data.NewGraphQL(c.gqlConfig)

	cty := city.City{
		Name: request.CityName,
		Lat:  request.Lat,
		Lng:  request.Lng,
	}

	cty, err = city.Add(context.Background(), gql, cty)
	if err != nil {
		c.log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := schema.UploadFeedResponse{
		CountryCode: "US",
		CityName:    cty.Name,
		Lat:         cty.Lat,
		Lng:         cty.Lng,
		Message:     "city added",
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
