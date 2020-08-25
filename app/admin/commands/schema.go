package commands

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/travel/business/data"
	"github.com/dgraph-io/travel/business/data/city"
	"github.com/dgraph-io/travel/business/data/schema"
	"github.com/dgraph-io/travel/business/data/weather"
	"github.com/pkg/errors"
)

// Schema handles the updating of the schema.
func Schema(gqlConfig data.GraphQLConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	gql := data.NewGraphQL(gqlConfig)

	schema, err := schema.New(gql)
	if err != nil {
		return errors.Wrapf(err, "preparing schema")
	}

	if err := schema.Create(ctx); err != nil {
		return errors.Wrapf(err, "creating schema")
	}

	return nil
}

// Seed handles loading the databse with city data.
func Seed(log *log.Logger, gqlConfig data.GraphQLConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	gql := data.NewGraphQL(gqlConfig)

	cty := city.City{
		Name: "bill",
		Lat:  34.87546,
		Lng:  -80.774834,
	}

	cty, err := city.Add(ctx, gql, cty)
	if err != nil {
		return err
	}

	wth := weather.Weather{
		CityName: "bill",
		Desc:     "just some weather",
		Temp:     98.7,
		MaxTemp:  99.4,
		MinTemp:  87.3,
		Pressure: 1098,
	}

	wth, err = weather.Replace(ctx, gql, cty.ID, wth)
	if err != nil {
		return err
	}

	fmt.Println("data seeded", cty.ID)
	return nil
}
