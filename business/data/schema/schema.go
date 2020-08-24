// Package schema provides schema support for the database.
package schema

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
)

// Schema error variables.
var (
	ErrNoSchemaExists = errors.New("no schema exists")
	ErrInvalidSchema  = errors.New("schema doesn't match")
)

// Schema provides support for schema operations against the database.
type Schema struct {
	graphql  *graphql.GraphQL
	document string
}

// New constructs a Schema value for use to manage the schema.
func New(graphql *graphql.GraphQL) (*Schema, error) {
	schema := Schema{
		graphql:  graphql,
		document: _document,
	}

	return &schema, nil
}

// Create is used create the schema in the database.
func (s *Schema) Create(ctx context.Context) error {
	schema, err := s.retrieve(ctx)
	if err != nil {
		return errors.Wrap(err, "can't create schema, db not ready")
	}

	// If the schema matches against what we know the
	// schema to be, don't try to update it.
	if err := s.validate(ctx, schema); err == nil {
		return nil
	}

	query := `mutation updateGQLSchema($schema: String!) {
		updateGQLSchema(input: {
			set: { schema: $schema }
		}) {
			gqlSchema {
				schema
			}
		}
	}`

	vars := map[string]interface{}{"schema": s.document}

	if err := s.graphql.QueryWithVars(ctx, graphql.CmdAdmin, query, vars, nil); err != nil {
		return errors.Wrap(err, "create schema")
	}

	schema, err = s.retrieve(ctx)
	if err != nil {
		return errors.Wrap(err, "can't create schema, db not ready")
	}

	if err := s.validate(ctx, schema); err != nil {
		return errors.Wrap(err, "invalid schema")
	}

	return nil
}

// retrieve queries the database for the schema and handles situations
// when the database is not ready for schema operations.
func (s *Schema) retrieve(ctx context.Context) (string, error) {
	for {
		schema, err := s.query(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "Server not ready") {

				// If the context deadline exceeded then we are done trying.
				if ctx.Err() != nil {
					return "", errors.Wrap(err, "server not ready")
				}

				// We need to wait for the server to be ready for this :(.
				time.Sleep(2 * time.Second)
				continue
			}
		}

		return schema, nil
	}
}

func (s *Schema) query(ctx context.Context) (string, error) {
	query := `query { getGQLSchema { schema }}`
	result := make(map[string]interface{})
	if err := s.graphql.QueryWithVars(ctx, graphql.CmdAdmin, query, nil, &result); err != nil {
		return "", errors.Wrap(err, "query schema")
	}

	data, err := json.Marshal(result)
	if err != nil {
		return "", errors.Wrap(err, "marshal schema")
	}

	return string(data), nil
}

func (s *Schema) validate(ctx context.Context, schema string) error {
	if schema == `{"getGQLSchema":null}` || schema == `{"getGQLSchema":{"schema":""}}` {
		return ErrNoSchemaExists
	}

	if len(schema) < 27 {
		return ErrInvalidSchema
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return errors.Wrap(err, "regex compile")
	}

	exp := strings.ReplaceAll(s.document, "\\n", "")
	exp = reg.ReplaceAllString(exp, "")
	schema = strings.ReplaceAll(schema[27:], "\\n", "")
	schema = strings.ReplaceAll(schema, "\\t", "")
	schema = reg.ReplaceAllString(schema, "")

	if exp != schema {
		return ErrInvalidSchema
	}

	return nil
}
