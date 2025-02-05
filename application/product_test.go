package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product()
	product.Name = "test"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Enable()
	require.nil(t, err)

	product.Price = 0
	err := product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product()
	product.Name = "test"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.nil(t, err)

	product.Price = 10
	err := product.Disable()
	require.Equal(t, "the price must be zero in order to disable the product", err.Error())
}
