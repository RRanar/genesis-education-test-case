package helpers

import (
	"genesis-education-test-case/core/services/storage"
)

func GetEmails(subs *[]storage.Subscriber) []string {
	var emails []string

	for _, v := range *subs {
		emails = append(emails, v.Email)
	}

	return emails
}