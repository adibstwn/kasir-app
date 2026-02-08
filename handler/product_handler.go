package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kasir-app/dto"
	"kasir-app/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.IProductService
}

func NewProductHandler(productService service.IProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (p *ProductHandler) GetAllProduct(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	maxRow, _ := strconv.Atoi(c.DefaultQuery("maxRow", "10"))

	productName := c.Query("productName")

	products, err := p.productService.GetAllProduct(productName, page, maxRow)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductHandler) GetProduct(c *gin.Context) {
	productID := c.Param("id")
	product, err := p.productService.GetById(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}

}

func (p *ProductHandler) CreateProduct(c *gin.Context) {
	product := &dto.CreateProduct{}
	err := c.BindJSON(product)
	if err != nil {
		log.Println("error binding createProduct json ", err)
		fmt.Println(err)
	}

	err = p.productService.Create(product)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
	})
}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	product := &dto.UpdateProduct{}
	err := c.BindJSON(product)
	if err != nil {
		log.Println("error binding updateProduct json ", err)
		fmt.Println(err)
	}

	err = p.productService.Update(product)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "updated",
		})
	}
}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	productID := c.Param("id")
	err := p.productService.Delete(productID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "deleted",
		})
	}
}
