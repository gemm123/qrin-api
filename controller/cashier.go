package controller

import (
	"gemm123/qrin-api/helper"
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controllerCashier struct {
	serviceCashier service.ServiceCashier
}

func NewControllerCashier(serviceCashier service.ServiceCashier) *controllerCashier {
	return &controllerCashier{serviceCashier: serviceCashier}
}

func (ctr *controllerCashier) Register(c *gin.Context) {
	var cashier models.Cashier
	if err := c.ShouldBindJSON(&cashier); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	passwordHash, err := helper.PasswordHash(cashier.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed " + err.Error(),
		})
		return
	}
	cashier.Password = passwordHash

	_, err = ctr.serviceCashier.Register(cashier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}
