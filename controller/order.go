package controller

import (
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type controllerOrder struct {
	serviceOrder   service.ServiceOrder
	serviceItem    service.ServiceItem
	serviceCashier service.ServiceCashier
}

func NewControllerOrder(serviceOrder service.ServiceOrder, serviceItem service.ServiceItem, serviceCashier service.ServiceCashier) *controllerOrder {
	return &controllerOrder{serviceOrder: serviceOrder, serviceItem: serviceItem, serviceCashier: serviceCashier}
}

func (ctr *controllerOrder) AddOrder(c *gin.Context) {
	var inputOrder models.InputOrder
	if err := c.ShouldBindJSON(&inputOrder); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	userID := c.MustGet("userID").(int)

	var layoutFormat, value string
	var date time.Time
	layoutFormat = "2006-01-02"
	value = inputOrder.Date
	date, _ = time.Parse(layoutFormat, value)

	cashier, err := ctr.serviceCashier.GetCashierByID(inputOrder.CashierID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	order := models.Order{
		ID:         inputOrder.ID,
		CashierID:  inputOrder.CashierID,
		UserID:     uint(userID),
		TotalPrice: inputOrder.TotalPrice,
		Payment:    inputOrder.Payment,
		Date:       date,
		Address:    cashier.Address,
		Shop:       cashier.Shop,
	}

	_, err = ctr.serviceOrder.AddOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	for _, inputItem := range inputOrder.Items {
		item, err := ctr.serviceItem.ShowDetailItemByName(inputItem.Name, inputOrder.CashierID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "faield " + err.Error(),
			})
			return
		}

		detailOrder := models.DetailOrder{
			Quantity: int64(inputItem.Quantity),
			Price:    inputItem.Price,
			OrderID:  inputOrder.ID,
			ItemID:   item.ID,
		}

		_, err = ctr.serviceOrder.AddDetailOrder(detailOrder)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "faield " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success add order",
	})
}

func (ctr *controllerOrder) ShowAllOrder(c *gin.Context) {
	userID := c.MustGet("userID").(int)
	order, err := ctr.serviceOrder.ShowAllOrder(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    order,
	})
}
