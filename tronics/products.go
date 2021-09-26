package tronics

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var products = []map[int]string{{1: "Watch"}, {2: "Phone"}, {3: "TV"}}

func getProduct(c echo.Context) error {
	var product map[int]string
	for _, p := range products {
		for j := range p {
			pId, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				return err
			}
			if pId == j {
				product = p
			}
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "No product found")
	}
	return c.JSON(http.StatusOK, product)
}
func getAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func addProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name"`
	}
	var reqBody body
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	product := make(map[int]string)
	product[len(products)+1] = reqBody.Name
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}
