package handler

import (
	"errors"
	"kasir-app/database"
	"kasir-app/model"
	"kasir-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	var user model.User
	if err := database.GormDB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"errors": "invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "failed generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

func Register(c *gin.Context) {
	var user model.User
	c.BindJSON(user)

	//return string(bytes), err

}

func x() {
	if err := y(0, 1); err != nil {

	}
}

func y(a int, b int) error {
	err := a % b
	if err < 1 {
		return errors.New("dibawah 1")
	}
	return nil
}
