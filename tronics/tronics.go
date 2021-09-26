package tronics

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		e.Logger.Fatal("Eroor")
	}
}

func Start() {
	e.GET("/products/:id", getProduct)
	e.GET("/products", getAllProducts)
	e.POST("/product", addProduct)
	e.Start(cfg.Port)
}
