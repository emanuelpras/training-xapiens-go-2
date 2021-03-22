package types

import "github.com/graphql-go/graphql"

var VendorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Vendor",
		Fields: graphql.Fields{
			"vendorID": &graphql.Field{
				Type: graphql.String,
			},
			"vendorName": &graphql.Field{
				Type: graphql.String,
			},
			"zipCode": &graphql.Field{
				Type: graphql.String,
			},
			"employee": &graphql.Field{
				Type: graphql.NewList(EmployeeType), //nested -> bikin list baru
			},
		},
	},
)
