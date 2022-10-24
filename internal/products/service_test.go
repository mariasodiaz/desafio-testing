package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	Products []Product
}

func (m MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	var products []Product
	for _, key := range m.Products {
		if key.SellerID == sellerID {
			products = append(products, key)
		}
	}
	if len(products) > 0 {
		return products, nil
	}
	return nil, errors.New("no se ha encontrado el id")
}

func TestGetAll(t *testing.T) {
	products := []Product{{ID: "1", SellerID: "123", Description: "una descripcion", Price: 1000.5}, {ID: "2", SellerID: "AC234", Description: "otra descripcion", Price: 100.12}}
	mock := MockRepository{Products: products}
	service := NewService(mock)
	expected := []Product{{ID: "1", SellerID: "123", Description: "una descripcion", Price: 1000.5}}
	result, err := service.GetAllBySeller("123")
	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	result, err = service.GetAllBySeller("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
