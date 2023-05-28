package rate

import (
	"genesis-education-test-case/core/services/rate"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRateController interface {
	New() *RateController
	GetRate(*gin.Context)
}

type RateController struct {
	rateService *rate.RateService
}

func New(apiKey string) *RateController {
	return &RateController{
		rateService: rate.New(apiKey),
	}
}

func (rateController *RateController) GetRate(c *gin.Context) {
	rate, err := rateController.rateService.GetRate()

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, rate)
}