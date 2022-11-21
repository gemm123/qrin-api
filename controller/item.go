package controller

import (
	"fmt"
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

	file, _ := c.FormFile("image")
	filename := strings.ReplaceAll(file.Filename, " ", "-")
	file.Filename = fmt.Sprintf("%d-%s", item.ID, filename)
	c.SaveUploadedFile(file, "src/images/"+file.Filename)

	filePath := "/images/" + file.Filename
	item.Image = filePath

	userID := c.MustGet("userID").(int)
	item.CashierID = uint(userID)

	newItem, err := ctr.serviceItem.AddItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": newItem,
	})
}
