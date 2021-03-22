package resolver

import (
	"log"
	"xapiens-project-gin-graphql/models"

	"github.com/graphql-go/graphql"
)

func UpdateProductResolve(p graphql.ResolveParams) (interface{}, error) {
	// ambil id nya terlebih dahulu, masukan ke dalam sebuah variable
	id := p.Args["id"].(int)
	name, checkName := p.Args["name"].(string)
	info, checkInfo := p.Args["info"].(string)
	price, checkPrice := p.Args["price"].(float64)

	log.Println("ini argsnyaa......", p.Args["name"].(string))

	productVar := models.Product{}

	for i, v := range models.ProductData { // product adalah database
		if id == int(v.ID) { // ketika id di database sama dengan args
			if checkName {
				models.ProductData[i].Name = name
			}
			if checkInfo {
				models.ProductData[i].Info = info
			}
			if checkPrice {
				models.ProductData[i].Price = price
			}
			productVar = models.ProductData[i]
			break
		}
	}
	return productVar, nil // untuk response yang akan ditampilkan
}
