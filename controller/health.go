package controller

import (
	"github.com/edwintantawi/go-todo-api/entity/health"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type healthController struct {
	service health.Service
}

func NewHealthController(s health.Service) *healthController {
	return &healthController{service: s}
}

func (c *healthController) HealthCheck(ctx *gin.Context) {
	status := http.StatusOK

	if err := c.service.Check(); err != nil {
		status = http.StatusServiceUnavailable
		log.Println(err)
	}

	ctx.JSON(status, gin.H{"message": http.StatusText(status)})
}
