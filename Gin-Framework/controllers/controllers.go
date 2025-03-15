package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string
	Name  string
	Price int
}

func ProductDummy() []Product {
	prod := []Product{
		{"c1", "Television", 100},
		{"c2", "Icebox", 200},
		{"c3", "Microwave", 10},
		{"c4", "Fan", 500},
	}
	return prod
}

var products = ProductDummy()

// GET
func GetProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"product": products,
	})
}

// POST
func CreateProduct(ctx *gin.Context) {
	var NewProduct Product

	// Parsing
	err := ctx.ShouldBindJSON(&NewProduct)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	NewProduct.ID = fmt.Sprintf("c%d", len(products)+1)
	products = append(products, NewProduct)

	ctx.JSON(http.StatusCreated, gin.H{
		"product": NewProduct,
	})
}

// Update
func UpdateProduct(ctx *gin.Context) {
	var UpdateProduct Product
	condition := false
	productID := ctx.Param("prodID")

	err := ctx.ShouldBindJSON(&UpdateProduct)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, prod := range products {
		if prod.ID == productID {
			condition = true
			products[i] = UpdateProduct
			products[i].ID = productID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"product": UpdateProduct,
	})
}

// Delete
func DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("prodID")
	condition := false
	var prodIndex int
	var deleteProduct Product 

	for i, prod := range products {
		if prod.ID == productID {
			condition = true
			prodIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":       "Product not found",
			"error_message": "Invalid product ID",
		})
		return
	}

	// Menghapus elemen dari slice secara aman
	deleteProduct = products[prodIndex]
	products = append(products[:prodIndex], products[prodIndex+1:]...)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
		"product" : deleteProduct,
	})
}
