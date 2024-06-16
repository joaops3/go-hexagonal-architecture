package application_test

import (
	"testing"

	"github.com/joaops3/go-exagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)


func TestProduct_Enable(t *testing.T){
	product := application.NewProduct("teste", 10)
	product.Price = 10
	product.Status = application.DISABLED
	err := product.Enable()

	require.Nil(t,err)
	require.Equal(t, application.ENABLED, product.GetStatus() )

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "o preço deve ser maior que 0 para habilitar o produto", err.Error() )
}

func TestProduct_Disable(t *testing.T){
	product := application.NewProduct("teste", 10)
	product.Status = application.DISABLED

	product.Price = 10
	err := product.Disable()
	require.Equal(t, "o preço deve ser 0 para desabilitar o produto", err.Error() )
}


func TestProduct_IsValid(t *testing.T){
	product := application.NewProduct("teste", 10)
	product.Price = 10
	product.Status = application.DISABLED
	product.Name = "Test"
	product.ID = uuid.NewV4().String()


	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "invalido"

	_, err = product.IsValid()

	require.Equal(t, "o status deve ser ativo o desativo", err.Error())

	
	product.ID = "123"

	_, err = product.IsValid()

	require.NotNil(t, err)

}