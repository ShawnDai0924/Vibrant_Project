package gql

import (
	"STATE-SEARCH-API/data"
	"context"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define the GraphQL State type
var stateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "State",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.ID},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Define the GraphQL Query type
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"states": &graphql.Field{
			Type: graphql.NewList(stateType),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, _ := p.Args["name"].(string)

				// Get state data from the database
				collection := p.Context.Value("mongoClient").(*mongo.Client).Database("mango").Collection("states")

				// Query all state data
				var results []data.State
				cursor, err := collection.Find(context.TODO(), bson.M{})
				if err != nil {
					return nil, err
				}
				defer cursor.Close(context.TODO())

				for cursor.Next(context.TODO()) {
					var state data.State
					if err := cursor.Decode(&state); err != nil {
						return nil, err
					}
					if name == "" || contains(state.Name, name) {
						results = append(results, state)
					}
				}

				return results, nil
			},
		},
	},
})

// Check if the name contains the keyword
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}

// Create a GraphQL Schema
func NewSchema(client *mongo.Client) (*graphql.Schema, error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		return nil, err
	}
	return &schema, nil
}
