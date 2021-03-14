package sentryLogger

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func Sentry(paramError error) {
	var DsnSentry string

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(err.Error())
	} else { // tidak ada error saat ambil file .env
		DsnSentry = os.Getenv("DSN_SENTRY") // isi value dari variable DsnSentry
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: DsnSentry,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureException(paramError)
}
