package cli

import (
	"fmt"

	"github.com/joaops3/go-exagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64)(string , error) {

	result := ""
	
	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("")
	}

	return "", nil
}