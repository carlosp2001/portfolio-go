package service

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"os"
	_struct "portfolio/backend/struct"
)

func SendEmail(form _struct.ContactForm) error {
	to := os.Getenv("SMTP_TO")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	client := resend.NewClient(smtpPassword)
	params := resend.SendEmailRequest{
		From:    "Portfolio Website <onboarding@resend.dev>",
		To:      []string{to},
		Subject: "New message from website",
		Html:    "<h2>New message from website</h1><p>" + form.Message + "</p>",
	}

	sent, err := client.Emails.Send(&params)
	if err != nil {
		return fmt.Errorf("Error sending the email: %w", err)
	}

	fmt.Println("Email ID:", sent.Id)
	return nil
}
