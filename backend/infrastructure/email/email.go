package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendVerificationEmail(to string, token string) error {
	// 環境変数の値をログ出力とチェック
	from := os.Getenv("EMAIL_ADDRESS")
	if from == "" {
		return fmt.Errorf("EMAIL_ADDRESS is not set")
	}

	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		return fmt.Errorf("SMTP_HOST is not set")
	}

	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		return fmt.Errorf("SMTP_PORT is not set")
	}

	log.Printf("Sending email from: %s", from)
	log.Printf("SMTP Settings - Host: %s, Port: %s", smtpHost, smtpPort)

	msg := []byte(fmt.Sprintf("From: <%s>\r\n"+
		"To: <%s>\r\n"+
		"Subject: Email Verification\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n"+
		"Please verify your email by clicking the following link: http://localhost:8080/verify?token=%s\r\n",
		from, to, token))

	// デバッグ用にメッセージを出力
	log.Printf("Message content:\n%s", string(msg))

	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, msg)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully to %s", to)
	return nil
}
