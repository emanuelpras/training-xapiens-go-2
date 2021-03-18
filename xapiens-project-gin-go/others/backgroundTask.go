package main

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func main() {
	cronJob()
}

func cronJob() {
	gocron.Every(6).Seconds().Do(taskWithoutParam)
	gocron.Every(30).Seconds().Do(SendEmail)
	gocron.Every(3).Seconds().Do(taskWithParam, 10, "sepuluh", true)
	<-gocron.Start()
}

func taskWithoutParam() {
	log.Println("Task without param")
}

func taskWithParam(param1 int, param2 string, param3 bool) {
	log.Println("Task without param", param1, "|", param2, "|", param3)
}

func SendEmailConfig(to []string, cc []string, subject, message string) error {
	// configuration untuk kirim email
	// smptp_hostname, email, password, smptp_port

	err := godotenv.Load(".env")

	if err == nil { // tidak ada error saat get file .env

		body := "From: " + os.Getenv("email_email") + "\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Cc : " + strings.Join(cc, ",") + "\n" +
			"Subject: " + subject + "\n\n" +
			message

		auth := smtp.PlainAuth("", os.Getenv("email_email"), os.Getenv("email_password"), os.Getenv("email_smptp_hostname"))
		smtpAdd := os.Getenv("email_smptp_hostname") + ":" + os.Getenv("email_smptp_port")
		err := smtp.SendMail(smtpAdd, auth, os.Getenv("email_email"), append(to, cc...),
			[]byte(body))

		if err == nil {
			return nil
		}
		return err
	}
	return err
}

func SendEmail() {
	to := []string{"sahlan.nasution07@gmail.com", "andityadimas@gmail.com"}
	cc := []string{"emanuelpras4@gmail.com"}

	subject := "Test email notification"
	message := "Halo gaes, ini tes email notif ya"

	err := SendEmailConfig(to, cc, subject, message)
	if err != nil { // kalau ada error saat send email
		log.Println(err.Error())
	}

	log.Println("Mail send!")
}
