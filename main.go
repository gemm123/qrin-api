package main

import (
	"fmt"
	"gemm123/qrin-api/config"
	"gemm123/qrin-api/controller"
	"gemm123/qrin-api/repository"
	"gemm123/qrin-api/service"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't get .env")
	}

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	config.InitDB(dsn)
	config.MirgrateDB()
	defer config.CloseDB()

	userRepository := repository.NewRepository(config.DB)
	userService := service.NewService(userRepository)
	userController := controller.NewController(userService)

	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	auth.POST("/register", userController.Register)

	r.Run(":" + port)
}
