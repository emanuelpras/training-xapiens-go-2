package api

import (
	"math/rand"
	"time"
	"xapiens-project-gin-graphql/models"
	"xapiens-project-gin-graphql/resolver"
	"xapiens-project-gin-graphql/types"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			// localhost:8080/?Query=mutation+_{create(name:"sprite",info:"minuman soda",price:5000){id,name,info,price}}
			"create": &graphql.Field{
				Type: types.ProductType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					products := models.Product{
						ID:    int64(rand.Intn(1000)),
						Name:  p.Args["name"].(string),
						Info:  p.Args["info"].(string),
						Price: p.Args["price"].(float64),
					}
					models.ProductData = append(models.ProductData, products)
					return products, nil
				},
			},

			"update": &graphql.Field{
				Type: types.ProductType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: resolver.UpdateProductResolve,
			},

			"delete": &graphql.Field{
				Type: types.ProductType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
					id := p.Args["id"].(int)

					productVar := models.Product{}

					for i, v := range models.ProductData { // product adalah database
						if id == int(v.ID) { // kalau data id dari database == data id dari args
							productVar = models.ProductData[i]
							models.ProductData = append(models.ProductData[:i], models.ProductData[i+1:]...)
						}
					}
					return productVar, nil // untuk response yang akan ditampilkan
				},
			},
		},
	},
)
