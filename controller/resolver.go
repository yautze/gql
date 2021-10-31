package controller

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"gql/controller/generated"
	"gql/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
