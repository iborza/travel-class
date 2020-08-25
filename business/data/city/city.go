// Package city provides support for managing city data in the database.
package city

import (
	"context"
	"fmt"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// Set of error variables for CRUD operations.
var (
	ErrExists   = errors.New("city exists")
	ErrNotFound = errors.New("city not found")
)

// Add adds a new city to the database. If the city already exists
// this function will fail but the found city is returned. If the city is
// being added, the city with the id from the database is returned.
func Add(ctx context.Context, gql *graphql.GraphQL, city City) (City, error) {
	city, err := add(ctx, gql, city)
	if err != nil {
		return City{}, errors.Wrap(err, "adding city to database")
	}

	return city, nil
}

// =============================================================================

func add(ctx context.Context, gql *graphql.GraphQL, city City) (City, error) {
	if city.ID != "" {
		return City{}, errors.New("city contains id")
	}

	mutation, result := prepareAdd(city)
	if err := gql.Query(ctx, mutation, &result); err != nil {
		return City{}, errors.Wrap(err, "failed to add city")
	}

	if len(result.AddCity.City) != 1 {
		return City{}, errors.New("city id not returned")
	}

	city.ID = result.AddCity.City[0].ID
	return city, nil
}

// =============================================================================

func prepareAdd(city City) (string, addResult) {
	var result addResult
	mutation := fmt.Sprintf(`
	mutation {
		addCity(input: [{
			name: %q
			lat: %f
			lng: %f
		}])
		%s
	}`, city.Name, city.Lat, city.Lng, result.document())

	return mutation, result
}
