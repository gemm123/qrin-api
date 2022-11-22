package controller

import (
	"fmt"
	"gemm123/qrin-api/helper"
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type controllerItem struct {
	serviceItem service.ServiceItem
}

func NewControllerItem(serviceItem service.ServiceItem) *controllerItem {
	return &controllerItem{serviceItem: serviceItem}
}

func (ctr *controllerItem) AddItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBind(&item); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	userID := c.MustGet("userID").(int)
	file, _ := c.FormFile("image")
	filename := strings.ReplaceAll(file.Filename, " ", "-")
	uniqueCode := helper.GenerateCodeImage()
	file.Filename = fmt.Sprintf("%d-%d-%s", userID, uniqueCode, filename)
	c.SaveUploadedFile(file, "src/images/"+file.Filename)

	filePath := "/images/" + file.Filename
	item.Image = filePath
	item.CashierID = uint(userID)

	newItem, err := ctr.serviceItem.AddItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    newItem,
	})
}

func (ctr *controllerItem) ShowAllItem(c *gin.Context) {
	cashierID := c.MustGet("userID").(int)
	items, err := ctr.serviceItem.ShowAllItem(uint(cashierID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    items,
	})
}
