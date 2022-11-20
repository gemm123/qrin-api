package controller

import (
	"gemm123/qrin-api/helper"
	"gemm123/qrin-api/models"
	"gemm123/qrin-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{service: service}
}

func (ctr *controller) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "faield " + err.Error(),
		})
		return
	}

	otp := helper.GenerateOTP()
	user.OTP = otp

	// request := gorequest.New()
	// resp, body, _ := request.Post("http://45.77.34.32:8000/demo/send_message").
	// 	Set("Authorization", "Basic dXNtYW5ydWJpYW50b3JvcW9kcnFvZHJiZWV3b293YToyNjM3NmVkeXV3OWUwcmkzNDl1ZA==").
	// 	Send(`{"phone_no":"+6282237436363", "key":"db63f52c1a00d33cf143524083dd3ffd025d672e255cc688", "message": "test"}`).
	// 	End()
	// fmt.Println(resp)
	// fmt.Println(body)

	passwordHash, err := helper.PasswordHash(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed " + err.Error(),
		})
		return
	}
	user.Password = passwordHash

	_, err = ctr.service.Register(user)
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
