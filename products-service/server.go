package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kolan92/producsts-service/models"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) RegisterRouter(router *gin.Engine) {

	router.GET("/check", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"check": "ok"})
	})

	s.registerCustomerEndpoints(router)
	s.registerProductEndpoints(router)
	s.registerOrdersEndpoints(router)
}

func (s *Server) registerCustomerEndpoints(router *gin.Engine) {
	customer := router.Group("/customer")

	customer.GET("/", func(c *gin.Context) {
		var customers []models.Customer
		if err := s.db.Find(&customers).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, customers)
		}
	})
	customer.POST("/", func(c *gin.Context) {

		var customer models.Customer
		if err := c.BindJSON(&customer); err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if err := s.db.Create(&customer).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, customer)
		}
	})
	customer.GET("/:customerID", func(c *gin.Context) {

		var customer models.Customer
		if err := s.db.Find(&customer, c.Param("customerID")).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, customer)
		}
	})
	customer.PUT("/:customerID", func(c *gin.Context) {

		var customer models.Customer
		if err := c.BindJSON(&customer); err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if err := s.db.Save(customer).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, customer)
		}
	})
	customer.DELETE("/:customerID", func(c *gin.Context) {

		customerID := c.Param("customerID")
		req := s.db.Delete(models.Customer{}, "ID = ?", customerID)
		if err := req.Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else if req.RowsAffected == 0 {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusOK)
		}
	})
}

func (s *Server) registerProductEndpoints(router *gin.Engine) {
	product := router.Group("/product")

	product.GET("/", func(c *gin.Context) {

		var products []models.Product
		if err := s.db.Find(&products).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, products)
		}
	})
	product.POST("/", func(c *gin.Context) {

		var product models.Product
		if err := c.BindJSON(&product); err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if err := s.db.Create(&product).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, product)
		}
	})
	product.GET("/:productID", func(c *gin.Context) {

		var product models.Product
		if err := s.db.Find(&product, c.Param("productID")).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, product)
		}
	})
	product.PUT("/:productID", func(c *gin.Context) {

		var product models.Product
		if err := c.BindJSON(&product); err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if err := s.db.Save(product).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, product)
		}
	})
	product.DELETE("/:productID", func(c *gin.Context) {

		productID := c.Param("productID")
		req := s.db.Delete(models.Product{}, "ID = ?", productID)
		if err := req.Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else if req.RowsAffected == 0 {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusOK)

		}
	})

}

func (s *Server) registerOrdersEndpoints(router *gin.Engine) {
	order := router.Group("/order")
	order.GET("/", func(c *gin.Context) {

		var orders []models.Order
		if err := s.db.Preload("Customer").Preload("Products").Find(&orders).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, orders)
		}
	})
	order.POST("/", func(c *gin.Context) {

		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if order.Customer.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must specify user"})
		}
		if err := s.db.Find(&order.Customer, order.Customer.ID).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		for i, product := range order.Products {
			if product.ID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "must specify a product ID"})
			}
			if err := s.db.Find(&order.Products[i], product.ID).Error; err != nil {
				c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
			}
		}

		if err := s.db.Create(&order).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, order)
		}
	})

	order.GET("/:orderID", func(c *gin.Context) {

		var order models.Order
		if err := s.db.Preload("Customer").Preload("Products").Find(&order, c.Param("orderID")).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, order)
		}
	})

	order.PUT("/:orderID", func(c *gin.Context) {

		var order models.Order
		if err := c.BindJSON(&order); err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if err := s.db.Model(&order).Save(order).Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, order)
		}
	})

	order.DELETE("/:orderID", func(c *gin.Context) {

		orderID := c.Param("orderID")
		req := s.db.Delete(models.Order{}, "ID = ?", orderID)
		if err := req.Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else if req.RowsAffected == 0 {
			c.Status(http.StatusNotFound)
		} else {
			c.Status(http.StatusOK)

		}
	})

	order.POST("/:orderID/product", func(c *gin.Context) {

		tx := s.db.Begin()

		var order models.Order
		orderID := c.Param("orderID")
		if err := tx.Preload("Products").First(&order, orderID).Error; err != nil {
			_ = tx.Rollback()
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		const productIDParam = "productID"
		productID := c.Query(productIDParam)
		if productID == "" {
			_ = tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("missing query param %q", productIDParam)})
		}

		var addedProduct models.Product
		if err := tx.First(&addedProduct, productID).Error; err != nil {
			_ = tx.Rollback()
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		order.Products = append(order.Products, addedProduct)
		if err := tx.Save(&order).Error; err != nil {
			_ = tx.Rollback()
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		}

		if err := tx.Commit().Error; err != nil {
			c.JSON(errToStatusCode(err), gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, order)
		}
	})
}

func errToStatusCode(err error) int {
	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
