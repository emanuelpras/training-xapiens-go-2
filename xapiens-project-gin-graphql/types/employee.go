package types

import "github.com/graphql-go/graphql"

var EmployeeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"employeeID": &graphql.Field{
				Type: graphql.String,
			},
			"fullName": &graphql.Field{
				Type: graphql.String,
			},
			"position": &graphql.Field{
				Type: graphql.String,
			},
			"vendorID": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
