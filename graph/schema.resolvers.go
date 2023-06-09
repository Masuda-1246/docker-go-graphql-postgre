package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/Masuda-1246/docker-go-graphql-postgre/graph/model"
	"github.com/Masuda-1246/docker-go-graphql-postgre/loader"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//ランダムな数字の生成
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand),
		UserId: input.UserID,
	}
	r.DB.Create(&todo)
	return &todo, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	rand, _ := rand.Int(rand.Reader, big.NewInt(100))
	user := model.User{
		ID:   fmt.Sprintf("U%d", rand),
		Name: input.Name,
	}
	r.DB.Create(&user)
	return &user, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	r.DB.Find(&todos)
	return todos, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user := []*model.User{}
	r.DB.Find(&user)
	return user, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user, err := loader.LoadUser(ctx, obj.UserId)
  if err != nil {
    return nil, err
  }
	return user, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	todo, err := loader.LoadTodo(ctx, obj.ID)
  if err != nil {
      return nil, err
  }
	return todo, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
