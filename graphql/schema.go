package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/SthiraPs/DeliveryApp/services"
	"github.com/SthiraPs/DeliveryApp/models"
)

// Define UserType
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.Int},
			"firstName": &graphql.Field{Type: graphql.String},
			"lastName":  &graphql.Field{Type: graphql.String},
			"email":     &graphql.Field{Type: graphql.String},
			"phoneNo":   &graphql.Field{Type: graphql.String},
			"roleId":    &graphql.Field{Type: graphql.Int},
			"statusId":  &graphql.Field{Type: graphql.Int},
		},
	},
)

// Define the Query schema
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// Fetch All Users
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return services.GetAllUsersService()
			},
		},
		// Fetch a User by ID
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				return services.GetUserByIdService(id)
			},
		},
	},
})

// Define the Mutation schema
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"firstName": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"lastName":  &graphql.ArgumentConfig{Type: graphql.String},
				"email":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"phoneNo":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"roleId":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"password":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"statusId":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := models.User{
					FirstName: p.Args["firstName"].(string),
					LastName:  p.Args["lastName"].(string),
					Email:     p.Args["email"].(string),
					PhoneNo:   p.Args["phoneNo"].(string),
					RoleId:    uint(p.Args["roleId"].(int)),
					Password:  p.Args["password"].(string),
					StatusId:  uint(p.Args["statusId"].(int)),
				}
				err := services.CreateUserService(user)
				return user, err
			},
		},
	},
})

// Create Schema
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
