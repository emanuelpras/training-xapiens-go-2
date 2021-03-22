package main

import (
	"net/http"

	"xapiens-project-gin-graphql/config"
	"xapiens-project-gin-graphql/executions"
	"xapiens-project-gin-graphql/schema"
	"xapiens-project-gin-graphql/seeder"

	"xapiens-project-gin-graphql/migrator"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connection()
	migrator.Migrations(dbPG)
	seeder.VendorSeeder(dbPG)
	seeder.EmployeSeeder(dbPG)

	route := gin.Default()
	route.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		var query Query

		c.Bind(&query)
		result := executions.ExecuteQuery(query.Query, schema.Schema)

		c.JSON(http.StatusOK, result)
	})

	route.Run()
}
