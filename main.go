package main

import (
	"fmt"
	"github.com/edwintantawi/go-todo-api/controller"
	"github.com/edwintantawi/go-todo-api/persistence/postgres"
	"github.com/edwintantawi/go-todo-api/repository"
	"github.com/edwintantawi/go-todo-api/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	db := postgres.New()

	healthRepository := repository.NewHealthRepository(db)
	healthService := service.NewHealthService(healthRepository)
	healthController := controller.NewHealthController(healthService)

	r := gin.Default()

	api := r.Group("/api")

	api.GET("/healthcheck", healthController.HealthCheck)

	if err := r.Run(PORT); err != nil {
		log.Fatal(err)
	}
}
