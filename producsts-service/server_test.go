package main

import (
	"log"
	"os"
	"testing"

	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	"github.com/gin-gonic/gin"
	"github.com/kolan92/producsts-service/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() {
	db, err := gorm.Open(sqlite.Open(":memory:?cache=shared"))
	if err != nil {
		panic("can't open database")
	}

	if err := db.AutoMigrate(&models.Customer{}, &models.Order{}, &models.Product{}); err != nil {
		panic(err)
	}

	router := gin.Default()
	server := NewServer(db)

	server.RegisterRouter(router)
	unitTest.SetRouter(router)
	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	unitTest.SetLog(newLog)
}

func TestCheck(t *testing.T) {
	setup()
	resp := make(map[string]string)

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/check", "json", nil, &resp)
	assert.NoError(t, err)

	checkValue := resp["check"]
	assert.Equal(t, "ok", checkValue)
}

// Customers
func TestGetCustomer(t *testing.T) {
	setup()
	var customers []models.Customer

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/customer/", "json", nil, &customers)
	assert.NoError(t, err)

	assert.Empty(t, customers)
}

func TestPostCustomer(t *testing.T) {
	setup()
	name := "Test Customer"
	newCustomer := models.Customer{
		ID:   1,
		Name: &name,
	}
	createdCustomer := models.Customer{}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/customer/", "json", &newCustomer, &createdCustomer)
	assert.NoError(t, err)
	assert.Equal(t, newCustomer, createdCustomer)

	var customers []models.Customer

	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/customer/", "json", nil, &customers)
	assert.NoError(t, err)

	assert.NotEmpty(t, customers)
	assert.Equal(t, newCustomer, customers[0])
}

func TestGetCustomerId(t *testing.T) {
	setup()
	name := "Test Customer"
	newCustomer := models.Customer{
		ID:   1,
		Name: &name,
	}

	createdCustomer := models.Customer{}
	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/customer/", "json", &newCustomer, &createdCustomer)
	assert.NoError(t, err)

	foundCustomer := models.Customer{}
	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/customer/1", "json", nil, &foundCustomer)
	assert.NoError(t, err)

	assert.Equal(t, newCustomer, foundCustomer)
}

func TestPutCustomer(t *testing.T) {
	setup()
	name := "Test Customer"
	newCustomer := models.Customer{
		ID:   1,
		Name: &name,
	}

	customerResponse := models.Customer{}
	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/customer/", "json", &newCustomer, &customerResponse)
	assert.NoError(t, err)

	updatedName := "Better customer"
	updatedCustomer := models.Customer{
		ID:   1,
		Name: &updatedName,
	}
	err = unitTest.TestHandlerUnMarshalResp(utils.PUT, "/customer/1", "json", &updatedCustomer, &customerResponse)
	assert.NoError(t, err)

	foundCustomer := models.Customer{}
	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/customer/1", "json", nil, &foundCustomer)
	assert.NoError(t, err)

	assert.Equal(t, updatedCustomer, foundCustomer)
}

// Products

func TestGetProduct(t *testing.T) {
	setup()
	var products []models.Product

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/product/", "json", nil, &products)
	assert.NoError(t, err)

	assert.Empty(t, products)
}

func TestPostProduct(t *testing.T) {
	setup()
	name := "Test Product"
	newProduct := models.Product{
		ID:   1,
		Name: &name,
	}
	createdProduct := models.Product{}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/product/", "json", &newProduct, &createdProduct)
	assert.NoError(t, err)
	assert.Equal(t, newProduct, createdProduct)

	var products []models.Product

	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/product/", "json", nil, &products)
	assert.NoError(t, err)

	assert.NotEmpty(t, products)
	assert.Equal(t, newProduct, products[0])
}

func TestGetProductId(t *testing.T) {
	setup()
	name := "Test Product"
	newProduct := models.Product{
		ID:   1,
		Name: &name,
	}

	createdProduct := models.Product{}
	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/product/", "json", &newProduct, &createdProduct)
	assert.NoError(t, err)

	foundProduct := models.Product{}
	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/product/1", "json", nil, &foundProduct)
	assert.NoError(t, err)

	assert.Equal(t, newProduct, foundProduct)
}

func TestPutProduct(t *testing.T) {
	setup()
	name := "Test Product"
	newProduct := models.Product{
		ID:   1,
		Name: &name,
	}

	productResponse := models.Product{}
	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/product/", "json", &newProduct, &productResponse)
	assert.NoError(t, err)

	updatedName := "Better product"
	updatedProduct := models.Product{
		ID:   1,
		Name: &updatedName,
	}
	err = unitTest.TestHandlerUnMarshalResp(utils.PUT, "/product/1", "json", &updatedProduct, &productResponse)
	assert.NoError(t, err)

	foundProduct := models.Product{}
	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/product/1", "json", nil, &foundProduct)
	assert.NoError(t, err)

	assert.Equal(t, updatedProduct, foundProduct)
}

// Orders

func TestGetOrder(t *testing.T) {
	setup()
	var products []models.Order

	err := unitTest.TestHandlerUnMarshalResp(utils.GET, "/product/", "json", nil, &products)
	assert.NoError(t, err)

	assert.Empty(t, products)
}

func TestPostOrder(t *testing.T) {
	setup()
	customerName := "Test Customer"
	newCustomer := models.Customer{
		ID:   1,
		Name: &customerName,
	}

	unitTest.TestHandlerUnMarshalResp(utils.POST, "/customer/", "json", &newCustomer, &models.Customer{})

	productName := "Test Product"
	newProduct := models.Product{
		ID:   1,
		Name: &productName,
	}

	unitTest.TestHandlerUnMarshalResp(utils.POST, "/product/", "json", &newProduct, &models.Product{})

	newOrder := models.Order{
		ID:       1,
		Customer: newCustomer,
		Products: []models.Product{newProduct},
	}
	createdOrder := models.Order{}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/order/", "json", &newOrder, &createdOrder)
	assert.NoError(t, err)
	assert.Equal(t, newOrder, createdOrder)

	var orders []models.Order

	err = unitTest.TestHandlerUnMarshalResp(utils.GET, "/order/", "json", nil, &orders)
	assert.NoError(t, err)

	assert.NotEmpty(t, orders)
	assert.Equal(t, newOrder, orders[0])
}
