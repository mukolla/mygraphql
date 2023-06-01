package schema

import (
	"github.com/graphql-go/graphql"
	"srvgraphql/pkg/resolver"
)

func CreateSchema() (graphql.Schema, error) {
	developerType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Developer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	workType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Work",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"startDate": &graphql.Field{
				Type: graphql.String,
			},
			"endDate": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	workDeveloperType := graphql.NewObject(graphql.ObjectConfig{
		Name: "WorkDeveloper",
		Fields: graphql.Fields{
			"developerId": &graphql.Field{
				Type: graphql.String,
			},
			"workId": &graphql.Field{
				Type: graphql.String,
			},
			"developer": &graphql.Field{
				Type:        developerType,
				Description: "The developer",
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"developer": &graphql.Field{
				Type:        developerType,
				Description: "Get a developer by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.ResolveDeveloper,
			},
			"developers": &graphql.Field{
				Type:        graphql.NewList(developerType),
				Description: "Get all developers",
				Resolve:     resolver.ResolveDevelopers,
			},
			"work": &graphql.Field{
				Type:        workType,
				Description: "Get a work by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "ID of the work",
					},
					"title": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Title of the work",
					},
				},
				Resolve: resolver.ResolveWork,
			},
			"works": &graphql.Field{
				Type:        graphql.NewList(workType),
				Description: "Get all developers",
				Resolve:     resolver.Works,
			},
			"workDeveloper": &graphql.Field{
				Type:        workDeveloperType,
				Description: "Get work developer relationship by developer ID and work ID",
				Args: graphql.FieldConfigArgument{
					"developerId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"workId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.ResolveWorkDeveloper,
			},
			"workDevelopers": &graphql.Field{
				Type:        graphql.NewList(workDeveloperType),
				Description: "Get all work developers",
				Resolve:     resolver.GetWorkDeveloper,
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"addDeveloper": &graphql.Field{
				Type:        developerType,
				Description: "Add a new developer",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					email, _ := p.Args["email"].(string)
					resolver.AddDeveloper(name, email)
					return resolver.Developers[len(resolver.Developers)-1], nil
				},
			},
		},
	})

	// Оголошуємо схему GraphQL
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}
