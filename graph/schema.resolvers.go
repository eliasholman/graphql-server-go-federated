package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"my/graphql/demo/graph/generated"
	"my/graphql/demo/graph/model"
)

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var product = &model.Product{
		ID: "2",
	}
	var order = &model.Order{
		ID:       "ABC1234",
		Products: []*model.Product{product},
		Buyer:    &model.Customer{ID: "2"},
	}
	return []*model.Order{order}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
