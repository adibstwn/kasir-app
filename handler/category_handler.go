package handler

import (
	"database/sql"
	"errors"
	"kasir-app/dto"
	"kasir-app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}
func (handler *CategoryHandler) GetAllCategory(c *gin.Context) {
	categories, _ := handler.categoryService.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
func (handler *CategoryHandler) GetCategoryById(c *gin.Context) {
	id := c.Param("id")
	category, err := handler.categoryService.GetById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "category not found",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": category,
		})
	}

}

func (handler *CategoryHandler) CreateCategory(c *gin.Context) {
	category := &dto.CreateCategoryReq{}
	err := c.BindJSON(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	err = handler.categoryService.Create(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Category Created",
		})
	}
}

func (handler *CategoryHandler) UpdateCategory(c *gin.Context) {
	category := &dto.UpdateCategoryReq{}
	err := c.BindJSON(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}
	err = handler.categoryService.Update(category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Category not found",
			})
		}
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Category Updated",
		})

	}
}
func (handler *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := handler.categoryService.DeleteById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Category not found",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Category Deleted",
		})
	}
}
