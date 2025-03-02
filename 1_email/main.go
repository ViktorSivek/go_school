package main

import (
	"fmt"
)

// Email struct to represent an email message
type Email struct {
	Sender    string
	Recipient string
	Message   string
}

// EmailService struct to manage emails
type EmailService struct {
	dbConnectionString string
}

// Method to send email based on protocol
func (e *EmailService) Send(email Email, protocol string) {
	switch protocol {
	case "SMTP":
		fmt.Println("Sending email via SMTP:", email)
	case "IMAP":
		fmt.Println("Sending email via IMAP:", email)
	case "POP3":
		fmt.Println("Sending email via POP3:", email)
	default:
		fmt.Println("Protocol not specified, using SMTP:", email)
	}
}

// Method to validate an email
func (e *EmailService) Validate(email Email) bool {
	if email.Sender == "" || email.Recipient == "" || email.Message == "" {
		fmt.Println("Invalid email: missing fields")
		return false
	}
	fmt.Println("Email is valid")
	return true
}

// Method to store email in a database
func (e *EmailService) Store(email Email) {
	fmt.Printf("Storing email from %s to %s in database\n", email.Sender, email.Recipient)
}

func main() {
	// Initialize EmailService
	emailService := EmailService{dbConnectionString: "db://email_service"}

	// Create an Email
	email := Email{Sender: "alice@example.com", Recipient: "bob@example.com", Message: "Hello Bob!"}

	// Validate and send the email
	if emailService.Validate(email) {
		emailService.Send(email, "SMTP")
		emailService.Store(email)
	}
}