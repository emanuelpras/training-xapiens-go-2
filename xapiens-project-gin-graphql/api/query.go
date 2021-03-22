package api

import (
	"xapiens-project-gin-graphql/config"
	"xapiens-project-gin-graphql/models"
	"xapiens-project-gin-graphql/types"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//localhost:8080/?Query={list{id, name, info, price}}
			"list": &graphql.Field{
				Type: graphql.NewList(types.ProductType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return models.ProductData, nil
				},
			},

			//localhost:8080/?Query={product(id:1){id, info, name, price}}
			"product": &graphql.Field{
				Type: types.ProductType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(param graphql.ResolveParams) (interface{}, error) {
					id, ok := param.Args["id"].(int)
					if ok {
						for _, products := range models.ProductData {
							if int(products.ID) == id {
								return products, nil
							}
						}
					}
					return nil, nil
				},
			},

			//localhost:8080/?Query={employee{employeeID, fullName, position}}
			"employee": &graphql.Field{
				Type: graphql.NewList(types.EmployeeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					dbPG := config.Connection()

					var (
						employee []models.Employee
					)
					dbPG.Find(&employee)
					return employee, nil
				},
			},
			//localhost:8080/?Query={vendor{vendorID, vendorName, zipCode, employee{employeeID, fullName, position, vendorID}}}
			"vendor": &graphql.Field{
				Type: graphql.NewList(types.VendorType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					dbPG := config.Connection()

					var (
						vendor []models.Vendor
					)
					dbPG.Preload("Employee").Find(&vendor)
					return vendor, nil
				},
			},
		},
	},
)
