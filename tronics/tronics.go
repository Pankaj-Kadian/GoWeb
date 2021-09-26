package tronics

import (
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func start() {
	e.GET("/products/:id", getProduct)
	e.GET("/products", getAllProducts)
	e.POST("/product", addProduct)
	e.Start(":8080")
}
