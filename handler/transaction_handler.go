package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kasir-app/dto"
	errors2 "kasir-app/errors"
	"kasir-app/service"
	"net/http"
)

type ITransactionHandler interface {
	Checkout(c *gin.Context)
	ReportToday(c *gin.Context)
	ReportByDate(c *gin.Context)
}

type TransactionHandler struct {
	transactionService service.ITransactionService
}

func NewTransactionHandler(transactionService service.ITransactionService) ITransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (t TransactionHandler) Checkout(c *gin.Context) {
	checkoutRequest := &dto.CheckoutRequest{}
	err := c.ShouldBindJSON(checkoutRequest)
	if err != nil {
		c.Error(errors2.BadRequest(err.Error()))
		return
	}

	err = t.transactionService.Checkout(c.Request.Context(), checkoutRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "transaction created",
	})
}

func (t TransactionHandler) ReportToday(c *gin.Context) {
	response, err := t.transactionService.ReportToday()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (t TransactionHandler) ReportByDate(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	response, err := t.transactionService.ReportByDate(startDate, endDate)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)

}
