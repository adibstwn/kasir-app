package handler

import (
	"github.com/gin-gonic/gin"
	"kasir-app/constant"
	"kasir-app/dto"
	"kasir-app/errors"
	"kasir-app/service"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) Create(c *gin.Context) {
	var userReq dto.CreateUserReq
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.Error(errors.BadRequest(err.Error()))
		return
	}

	err := handler.service.CreateUser(&userReq)
	if err != nil {
		c.Error(err)
		return
	}

	c.Set(constant.ResponseStatusCode, 201)
	c.Set(constant.ResponseMessage, "created")
}
