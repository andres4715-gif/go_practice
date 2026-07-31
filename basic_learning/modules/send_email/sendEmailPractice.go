package main

/*
This email is sent using an external Go module https://github.com/go-gomail/gomail
- The service https://mailtrap.io was also used
- With the link: https://mailtrap.io/inboxes/4438940/messages/5371965058
- Mailtrap account username and password were used
*/

import (
	"github.com/joho/godotenv" // Our new package
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. EXTRACT THE PASSWORDS INTO SECURE VARIABLES
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// 1. "Building" the shopping cart (the message)
	m := gomail.NewMessage()

	// Configure basic headers
	m.SetHeader("From", "andres4715@gmail.com")
	m.SetHeader("To", "andres4715@gmail.com	")
	m.SetHeader("Subject", "My second email from Go!")

	// Configure the email body (Supports HTML)
	m.SetBody("text/html", "<b>Congratulations!</b> You just sent your first email using an external module in Go 🚀.")

	// 2. Configure the "Mailman" (SMTP Server)
	// Parameters: Host, Port, User, Password
	dialer := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io", // Host
		2525,                       // Port (no quotes because it's a number)
		smtpUser,                   // User from environment variable
		smtpPass,                   // Password from environment variable
	)
	// 3. Send the email and handle the error the Senior way
	log.Println("⏳ Sending email, please wait...")

	if err := dialer.DialAndSend(m); err != nil {
		log.Fatalf("🚨 Critical error sending email: %v", err)
	}

	// If we got here, the error was 'nil'
	log.Println("✅ Email sent successfully!")
}