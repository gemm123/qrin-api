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

func (ctr *controllerCashier) LoginCashier(c *gin.Context) {
	var input models.InputLoginCashier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	err := ctr.serviceCashier.CheckEmailCashier(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed wrong email or password",
		})
		return
	}

	ok, err := ctr.serviceCashier.CheckPassword(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed " + err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed wrong email or password",
		})
		return
	}

	cashier, err := ctr.serviceCashier.GetCashier(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed " + err.Error(),
		})
		return
	}

	signedToken, err := ctr.serviceCashier.GenerateToken(cashier.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token":   signedToken,
	})
}

func (ctr *controllerCashier) GetCashier(c *gin.Context) {
	userID := c.MustGet("userID").(int)

	cashier, err := ctr.serviceCashier.GetCashierByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": cashier,
	})
}
