package main

import (
	"fmt"
	"genesis-education-test-case/core/routes"
	emailsender "genesis-education-test-case/core/services/email_sender"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	f, err := os.OpenFile(os.Getenv("STORAGE_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		fmt.Printf("Couldn't open file %s", os.Getenv("STORAGE_FILE"))
		return
	}

	f.Close()
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		fmt.Println("Cannot convert \"SMTP_PORT\" to number")
		return
	}

	routes.Subscribe(
		&app.RouterGroup,
		f,
		&emailsender.EmailSenderConfig{
			SmtpUser: os.Getenv("SMTP_USERNAME"),
			SmtpPass: os.Getenv("SMTP_PASSWORD"),
			SmtpSender: os.Getenv("SMTP_SENDER"),
			SmtpHost: os.Getenv("SMTP_HOST"),
			SmtpPort: int16(smtpPort),
		},
		os.Getenv("EXCHANGE_API_KEY"),
	)
	routes.Rate(&app.RouterGroup, os.Getenv("EXCHANGE_API_KEY"))

	app.Run("localhost:" + string(os.Getenv("PORT")))
}
