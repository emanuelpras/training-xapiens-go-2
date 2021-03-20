package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// bikin struct untuk model (nantinya)
type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info"`
	Price float64 `json:"price"`
}

// inisialisasi data (macem databasenya)
var product = []Product{
	{
		ID:    1,
		Name:  "Kopi luwak",
		Info:  "Nyaman di lambung nggak bikin kembung",
		Price: 1500,
	},
	{
		ID:    2,
		Name:  "Indomie",
		Info:  "Mie paling mantap",
		Price: 2500,
	},
	{
		ID:    3,
		Name:  "Coki coki",
		Info:  "Coklat asli",
		Price: 500,
	},
}

// bikin variable supaya graphQL tau field yang ada di model yg akan dipakai
var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// query type itu kayak method get di rest
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// get all data product
			// url nya localhost:8080/?Query={list{id, name, info, price}}
			"list": &graphql.Field{
				Type: graphql.NewList(productType),
				// kalau di rest resolve itu kayak controller
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return product, nil
				},
			},

			// get data with param
			// url untuk akses endpointnya localhost:8080/?Query={product(id:1){id, info, name, price}}
			"product": &graphql.Field{
				Type: productType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: func(param graphql.ResolveParams) (interface{}, error) {
					id, ok := param.Args["id"].(int)
					if ok {
						for _, products := range product { // looping data dari slice product
							// lakukan pengecekan jika sama id di args == id di slice product
							// maka tampilkan data
							if int(products.ID) == id {
								return products, nil
							}
						}
					}
					return nil, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			// urlnya untuk create data
			// localhost:8080/?Query=mutation+_{create(name:"sprite",info:"minuman soda",price:5000){id,name,info,price}}
			"create": &graphql.Field{
				Type: productType,
				Args: graphql.FieldConfigArgument{ // untuk param
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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano()) // ini untuk generate random value

					// bikin variable untuk parse data ke struct Product
					products := Product{
						ID:    int64(rand.Intn(1000)),
						Name:  p.Args["name"].(string),
						Info:  p.Args["info"].(string),
						Price: p.Args["price"].(float64),
					}
					product = append(product, products) //ini untuk nambah data di database (var product)
					return products, nil                // untuk response yang akan ditampilkan
				},
			},

			// update data by id
			//
			"update": &graphql.Field{
				Type: productType,
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
				Resolve: updateProductResolve,
			},

			"delete": &graphql.Field{
				Type: productType,
				Args: graphql.FieldConfigArgument{ // untuk param
					"id": &graphql.ArgumentConfig{ // id nggak boleh kosong
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				// kalau di rest resolve itu kayak controller
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
					id := p.Args["id"].(int)

					productVar := Product{}

					for i, v := range product { // product adalah database
						if id == int(v.ID) { // kalau data id dari database == data id dari args
							productVar = product[i]
							// untuk hapus bisa menggunakan slice
							product = append(product[:i], product[i+1:]...) // ini untuk hapus data dengan penerapan slice
							log.Println("ini productnya....", product)
						}
					}
					return productVar, nil // untuk response yang akan ditampilkan
				},
			},
		},
	},
)

func updateProductResolve(p graphql.ResolveParams) (interface{}, error) {
	// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
	id := p.Args["id"].(int)
	name, checkName := p.Args["name"].(string)
	info, checkInfo := p.Args["info"].(string)
	price, checkPrice := p.Args["price"].(float64)

	log.Println("ini argsnyaa......", p.Args["name"].(string))

	productVar := Product{}

	for i, v := range product { // product adalah database
		if id == int(v.ID) { // ketika id di database sama dengan args
			if checkName {
				product[i].Name = name
			}
			if checkInfo {
				product[i].Info = info
			}
			if checkPrice {
				product[i].Price = price
			}
			productVar = product[i]
			break
		}
	}
	return productVar, nil // untuk response yang akan ditampilkan
}

// semacam routing di rest
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,    // query itu hanya untuk get data
		Mutation: mutationType, //ini untuk Create Update Delete
	},
)

// function untuk execute query
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// check error condition saat run var result
	if len(result.Errors) > 0 {
		fmt.Println("ada error : ", result.Errors)
	}

	return result
}

func main() {
	route := gin.Default()
	// routing "/" ibarat di rest seperti group
	route.PUT("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)
		result := executeQuery(query.Query, schema)

		c.JSON(http.StatusOK, result)
	})

	route.Run()
}
