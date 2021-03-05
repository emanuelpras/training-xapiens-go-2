// konfigurasi untuk databse
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//funcion untuk konfig koneksi
func Connection() *gorm.DB {

	/* userDB := "postgres"
	passDB := "Komponen1!"
	hostDB := "localhost"
	portDB := "5432"
	namaDB := "vendorPayload"
	ssl := "disable"
	timeZone := "Asia/Jakarta" */
	var userDB, passDB, hostDB, portDB, namaDB, ssl, timeZone string
	//if condition untuk check file .env
	//kalau nggak ada filenya, munculin error
	//kalau ada filenya kita set env variablenya
	if err := godotenv.Load(".env"); err != nil { //kalau file nggak ketemu
		log.Fatalf(err.Error())
	} else { // kalau file ketemu
		userDB = os.Getenv("DATABASE")
		passDB = os.Getenv("DATABASE_PASS")
		hostDB = os.Getenv("DATABASE_HOST")
		portDB = os.Getenv("DATABSE_PORT")
		namaDB = os.Getenv("DATABASE_NAME")
		ssl = os.Getenv("DATABASE_SSLMODE")
		timeZone = os.Getenv("DATABASE_TIMEZONE")
	}

	conn :=
		" host=" + hostDB +
			" user=" + userDB +
			" password=" + passDB +
			" dbname=" + namaDB +
			" port=" + portDB +
			" sslmode=" + ssl +
			" TimeZone=" + timeZone

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	// kondisi ketika akses database postgre
	if err != nil {
		//panic
		panic("Gagal masuk ke database")
	} else {
		fmt.Println("Koneksi ke database berhasil")
	}

	return db
}
