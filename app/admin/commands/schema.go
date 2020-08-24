package commands

import (
	"context"
	"time"

	"github.com/dgraph-io/travel/business/data"
	"github.com/dgraph-io/travel/business/data/schema"
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
