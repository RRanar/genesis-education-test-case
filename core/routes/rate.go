package routes

import (
	"github.com/gin-gonic/gin"
	"genesis-education-test-case/core/controllers/rate"
)

func Rate(g *gin.RouterGroup, rateApiKey string) {
	rateController := rate.New(rateApiKey)
	g.GET("/rate", rateController.GetRate)
}