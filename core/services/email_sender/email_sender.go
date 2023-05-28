package emailsender

import (
	"github.com/go-mail/mail"
)

const DEFAULT_EMAIL_TITLE = "Rate: BTC to UAH"

type EmailSender interface {
	New(*EmailSenderConfig) *EmailSenderService
	Send([]string, string) error
}

type EmailSenderConfig struct {
	SmtpUser string
	SmtpPass string
	SmtpSender string
	SmtpHost string
	SmtpPort int16
}

type EmailSenderService struct {
	smtpConfig *EmailSenderConfig
}

func New(smtpConfig *EmailSenderConfig) *EmailSenderService {
	return &EmailSenderService{
		smtpConfig: smtpConfig,
	}
}

func (emailSender *EmailSenderService) Send(addr []string, message string) error {
	m := mail.NewMessage()

	m.SetHeader("From", emailSender.smtpConfig.SmtpSender)

	m.SetHeader("To", addr...)

	m.SetHeader("Subject", DEFAULT_EMAIL_TITLE)

	m.SetBody("text/plain", message)

	d := mail.NewDialer(
		emailSender.smtpConfig.SmtpHost,
		int(emailSender.smtpConfig.SmtpPort),
		emailSender.smtpConfig.SmtpUser,
		emailSender.smtpConfig.SmtpPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}