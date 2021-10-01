package book

import (
	"github.com/graphql-go/graphql"
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"book": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"bookByName": &graphql.Field{
				Type:        productType,
				Description: "Get book by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					name, ok := p.Args["name"].(string)
					if ok {
						// Find product
						result = FetchBookByName(name)
					}
					return result, nil
				},
			},
			"book": &graphql.Field{
				Type:        productType,
				Description: "Get book by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						result = FetchBookById(id)
					}
					return result, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get book list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					limit, _ := params.Args["limit"].(int)
					result = FetchBookList(limit)
					return result, nil
				},
			},
			"AuthorByName": &graphql.Field{
				Type:        authorType,
				Description: "Get book by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					name, ok := p.Args["name"].(string)
					if ok {
						// Find product
						result = FetchAuthorByName(name)
					}
					return result, nil
				},
			},
			"author": &graphql.Field{
				Type:        authorType,
				Description: "Get Author by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						result = FetchAuthorById(id)
					}
					return result, nil
				},
			},
			"author_list": &graphql.Field{
				Type:        graphql.NewList(authorType),
				Description: "Get book list",
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var result interface{}
					limit, _ := params.Args["limit"].(int)
					result = FetchAuthorList(limit)
					return result, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createBook": &graphql.Field{
			Type:        productType,
			Description: "Create new book",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				book := BookEntity{
					Name:        params.Args["name"].(string),
					Author:       params.Args["author"].(string),
				}
				if err := InsertBookEntity(book); err != nil {
					return nil, err
				}

				return book, nil
			},
		},

		"updateBook": &graphql.Field{
			Type:        productType,
			Description: "Update book by name",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				book := BookEntity{}
				if name, nameOk := params.Args["name"].(string); nameOk {
					book.Name = name
				}
				if author, ok := params.Args["author"].(string); ok {
					book.Author = author
				}

				if err := UpdateBookEntity(book); err != nil {
					return nil, err
				}
				return book, nil
			},
		},

		"deleteBook": &graphql.Field{
			Type:        productType,
			Description: "Delete book by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				if err := DeleteBookEntityById(id); err != nil {
					return nil, err
				}
				return id, nil
			},
		},
		"createAuthor": &graphql.Field{
			Type:        authorType,
			Description: "Create new author",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"book": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				book := AuthorEntity{
					Name:        params.Args["name"].(string),
					Book:       params.Args["book"].(int),
				}
				if err := InsertAuthorEntity(book); err != nil {
					return nil, err
				}

				return book, nil
			},
		},

		"updateAuthor": &graphql.Field{
			Type:        authorType,
			Description: "Update book by name",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"author": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				author := AuthorEntity{}
				if name, nameOk := params.Args["name"].(string); nameOk {
					author.Name = name
				}
				if book, ok := params.Args["book"].(int); ok {
					author.Book = book
				}

				if err := UpdateAuthorEntity(author); err != nil {
					return nil, err
				}
				return author, nil
			},
		},

		"deleteAuthor": &graphql.Field{
			Type:        authorType,
			Description: "Delete book by name",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id, _ := params.Args["id"].(int)
				if err := DeleteAuthorEntityById(id); err != nil {
					return nil, err
				}
				return id, nil
			},
		},
	},
})

// schema
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)
