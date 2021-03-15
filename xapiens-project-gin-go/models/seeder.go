package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeederUser(db *gorm.DB) {
	var userArray = [...]string{
		"admin",
		"test",
		"guest",
	}

	var user Users

	for _, v := range userArray {
		user.Username = v
		user.Role = v
		user.Password = v

		// enkrip password
		// param 1 dari password yang sudah ditentukan (contoh : admin)
		// kita bikin konfersi dari string ke byte -> caranya []byte(user.Password)
		encrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		// checking apakah proses enkripsi error / tidak
		if err != nil {
			log.Println(err)
		}

		user.Password = string(encrypt)
		user.ID = 0 // declare id dimulai dari 0, karena auto increment
		db.Create(&user)
	}
	fmt.Println("Seeder user created")
}
