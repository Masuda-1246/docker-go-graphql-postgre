package graph

import "github.com/Masuda-1246/docker-go-graphql-postgre/graph/model"
//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
}
