package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Categories struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories []Categories

func RegisterRoutes(r *gin.Engine) {

	// simple route
	r.GET("/categories", getAllCategories)
	r.POST("/categories", createCategory)
	r.PUT("/category/:id", updateCategory)
	r.GET("/category/:id", getCategory)
	r.DELETE("/category/:id", deleteCategory)

	// route group
	//api := r.Group("/api")
	//{
	//	api.GET("/users", getUsers)
	//	api.POST("/users", createCategory)
	//}
}

func deleteCategory(c *gin.Context) {
	for i, category := range categories {
		if category.Id == c.Param("id") {
			categories = append(categories[:i], categories[i+1:]...)
			c.JSON(200, gin.H{
				"status": "category deleted",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "category not found",
	})
	return
}

func updateCategory(c *gin.Context) {
	var req CreateCategoryReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, category := range categories {
		if category.Id == c.Param("id") {
			categories[i] = Categories{
				Id:          c.Param("id"),
				Name:        req.Name,
				Description: req.Description,
			}
			c.JSON(200, gin.H{
				"data": categories[i],
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "category not found",
	})
	return
}

func getCategory(c *gin.Context) {
	for _, category := range categories {
		if category.Id == c.Param("id") {
			c.JSON(200, gin.H{
				"data": category,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error": "category not found",
	})
	return
}

func createCategory(c *gin.Context) {
	var req CreateCategoryReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, category := range categories {
		fmt.Println(category.Name + " before")
	}
	categories = append(categories, Categories{
		Id:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
	})

	for _, category := range categories {
		fmt.Println(category.Name + " after")
	}

	c.JSON(201, gin.H{
		"status": "category created",
	})
}

func getAllCategories(c *gin.Context) {
	c.JSON(200, gin.H{
		"categories": categories,
	})
}
