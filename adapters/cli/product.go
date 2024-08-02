package cli

import (
	"fmt"

	"github.com/joaops3/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64)(string , error) {

	result := ""
	
	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID %s with name %s has been created with price %f ans status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus() )
	case "enable":
	product, err := service.Create(productName, productPrice)
	if err != nil {
			return "", err
	}
	res, err := service.Enable(product)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product %s has been enabled", res.GetName())
	case "disabled":
	product, err := service.Create(productName, productPrice)
	if err != nil {
			return "", err
	}
	res, err := service.Disable(product)
	if err != nil {
		return result, err
	}
	result = fmt.Sprintf("Product %s has been disabled", res.GetName())

	default:
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\n name %s\n, price %f\n status %s\n", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus() )
	
	}
	
	return result, nil
}