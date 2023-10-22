package utils

import (
	"testing"

	"github.com/risqiikhsani/rentvehicles/configs"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {

	secretsPath := "../"
	secretConfig, err := configs.LoadSecretConfig(secretsPath)
	if err != nil {
		panic(err)
	}
	// Get environment variables or provide default values.
	email_sender_name := secretConfig.EmailSenderName
	email_sender_address := secretConfig.EmailSenderAddress
	email_sender_password := secretConfig.EmailSenderPassword

	// Initialize the email sender.
	sender := NewGmailSender(email_sender_name, email_sender_address, email_sender_password)

	subject := "A test email"
	content := `
    <h1>Hello world</h1>
    <p>This is a test message from <a href="http://techschool.guru">Tech School</a></p>
    `
	to := []string{"risqiikhsani12@gmail.com"}
	// attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	require.NoError(t, err)
}
