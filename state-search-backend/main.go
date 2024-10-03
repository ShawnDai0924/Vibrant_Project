package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"STATE-SEARCH-API/gql"

	"github.com/graphql-go/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Create and connect the MongoDB client
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			log.Fatalf("failed to disconnect from mongo: %v", err)
		}
	}()

	//  GraphQL Schema
	schema, err := gql.NewSchema(mongoClient)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// GraphQL handler
	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true, // Enable GraphiQL (for testing)
	})

	// CORS (fix for not fetching to states)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		ctx := context.WithValue(r.Context(), "mongoClient", mongoClient)
		r = r.WithContext(ctx)

		// Call the GraphQL handler
		h.ServeHTTP(w, r)
	})

	// Start the HTTP server
	fmt.Println("Now server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
