package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"my/graphql/demo/graph/generated"
	"my/graphql/demo/graph/model"
)

func (r *entityResolver) FindCustomerByID(ctx context.Context, id string) (*model.Customer, error) {
	return &model.Customer{ID: id}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
