// Package user provides support for managing users in the database.
package user

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/graphql"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Set of error variables for CRUD operations.
var (
	ErrNotExists = errors.New("user does not exist")
	ErrExists    = errors.New("user exists")
	ErrNotFound  = errors.New("user not found")
)

// Add adds a new user to the database. If the user already exists
// this function will fail but the found user is returned. If the user is
// being added, the user with the id from the database is returned.
func Add(ctx context.Context, gql *graphql.GraphQL, nu NewUser, now time.Time) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, errors.Wrap(err, "generating password hash")
	}

	u := User{
		Name:         nu.Name,
		Email:        nu.Email,
		Role:         nu.Role,
		PasswordHash: string(hash),
		DateCreated:  now,
		DateUpdated:  now,
	}

	u, err = add(ctx, gql, u)
	if err != nil {
		return User{}, errors.Wrap(err, "adding user to database")
	}

	return u, nil
}

// =============================================================================

func add(ctx context.Context, gql *graphql.GraphQL, user User) (User, error) {
	mutation, result := prepareAdd(user)
	if err := gql.Query(ctx, mutation, &result); err != nil {
		return User{}, errors.Wrap(err, "failed to add user")
	}

	if len(result.AddUser.User) != 1 {
		return User{}, errors.New("user id not returned")
	}

	user.ID = result.AddUser.User[0].ID
	return user, nil
}

// =============================================================================

func prepareAdd(user User) (string, addResult) {
	var result addResult
	mutation := fmt.Sprintf(`
mutation {
	addUser(input: [{
		name: %q
		email: %q
		role: %s
		password_hash: %q
		date_created: %q
		date_updated: %q
	}])
	%s
}`, user.Name, user.Email, user.Role, user.PasswordHash,
		user.DateCreated.UTC().Format(time.RFC3339),
		user.DateUpdated.UTC().Format(time.RFC3339),
		result.document())

	return mutation, result
}
