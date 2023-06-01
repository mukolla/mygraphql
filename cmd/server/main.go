package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	schema "srvgraphql/schema/graphql"
)

func main() {
	schema, err := schema.CreateSchema()
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {

		var reqBody struct {
			Query    string `json:"query"`
			Mutation string `json:"mutation"`
		}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		var result *graphql.Result
		if len(reqBody.Query) > 0 {
			result = executeQuery(reqBody.Query, schema)
		} else if len(reqBody.Mutation) > 0 {
			result = executeQuery(reqBody.Mutation, schema)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(result)
	})

	log.Println("GraphQL server started on http://localhost:8082/graphql")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {

	fmt.Println(query)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("GraphQL query execution errors: %v", result.Errors)
	}
	return result
}
