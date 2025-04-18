package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joaops3/go-hexagonal/adapters/cli"
	mock_application "github.com/joaops3/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "product Test"
	productPrice := 10.00
	productStatus := "ENABLED"
	productId := "123"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with name %s has been created with price %f ans status %s", productId, productName, productPrice, productStatus )
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)


	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)


	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)


	resultExpected = fmt.Sprintf("Product ID: %s\n name %s\n, price %f\n status %s\n", productId, productName, productPrice, productStatus )
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}