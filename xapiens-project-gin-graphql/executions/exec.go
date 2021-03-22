package executions

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
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
