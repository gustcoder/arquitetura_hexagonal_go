package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gustcoder/arquitetura_hexagonal_go/adapters/cli"
	mock_application "github.com/gustcoder/arquitetura_hexagonal_go/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "123"
	productName := "Rise of the Ronin"
	productPrice := 147.00
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultEnableExpected := fmt.Sprintf("Product %s has been enabled", productName)
	resultEnable, err := cli.Run(service, "enable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultEnableExpected, resultEnable)

	resultDisableExpected := fmt.Sprintf("Product %s has been disabled", productName)
	resultDisable, err := cli.Run(service, "disable", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultDisableExpected, resultDisable)

	resultDefaultExpected := fmt.Sprintf("Product ID: %s\n Name: %s\n Price: %f\n Status: %s",
		productId, productName, productPrice, productStatus)
	resultDefault, err := cli.Run(service, "default", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultDefaultExpected, resultDefault)
}
