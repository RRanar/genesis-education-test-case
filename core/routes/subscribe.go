package routes

import (
	"genesis-education-test-case/core/controllers/subscribe"
	emailsender "genesis-education-test-case/core/services/email_sender"
	"os"

	"github.com/gin-gonic/gin"
)

func Subscribe(g *gin.RouterGroup, file *os.File, emailConf *emailsender.EmailSenderConfig, rateApiKey string) {
	subController := subscriber.New(file, emailConf, rateApiKey)
	g.POST("/subscribe", subController.Subscribe)
	g.POST("/sendEmails", subController.SendMessages)
}