package subscriber

import (
	"fmt"
	emailsender "genesis-education-test-case/core/services/email_sender"
	"genesis-education-test-case/core/services/rate"
	"genesis-education-test-case/core/services/storage"
	"genesis-education-test-case/core/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SubscribeRequestBody struct {
	Email string	`form:"email"`
}

type ISubscriberController interface {
	New(*os.File, *emailsender.EmailSenderConfig, string) *SubscriberController
	Subscribe(*gin.Context)
	SendMessages(*gin.Context)
}

type SubscriberController struct {
	subsStorage *storage.SubscribeStorageService
	emailSender *emailsender.EmailSenderService
	rateService *rate.RateService
}

func New(file *os.File, emailConf *emailsender.EmailSenderConfig, rateApiKey string) *SubscriberController {
	return &SubscriberController{
		subsStorage: storage.New(file),
		emailSender: emailsender.New(emailConf),
		rateService: rate.New(rateApiKey),
	}
}

func (subController *SubscriberController) Subscribe(c *gin.Context) {
	var subReqBody SubscribeRequestBody

	if err := c.Bind(&subReqBody); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, nil)
		return
	}

	sub, err := subController.subsStorage.FindByEmail(subReqBody.Email)

	if sub != nil || err != nil {
		c.JSON(http.StatusConflict, nil)
		return
	}

	if err := subController.subsStorage.Save(storage.Subscriber{Email: subReqBody.Email}); err != nil {
		c.JSON(http.StatusConflict, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (subController *SubscriberController) SendMessages(c *gin.Context) {
	rate, err := subController.rateService.GetRate()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	subs, err := subController.subsStorage.FindAll()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if err := subController.emailSender.Send(helpers.GetEmails(subs), fmt.Sprintf("Curent BTC to UAH rate: %.2f", rate)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}