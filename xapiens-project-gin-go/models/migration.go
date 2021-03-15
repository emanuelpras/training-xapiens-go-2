package models

import (
	"fmt"

	"gorm.io/gorm"
)

//function untuk migrasi ke database
func Migrations(db *gorm.DB) {

	// drop table user
	db.Migrator().DropTable(&Users{})

	// checking apakah table yang dimaksud sudah ada di db postgre
	// kalau belum ada, otomatis migrate
	if check := db.Migrator().HasTable(&Vendor{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&Vendor{})
		fmt.Println("Table vendor berhasil di create")
	}

	if check := db.Migrator().HasTable(&Employee{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&Employee{})
		fmt.Println("Table employee berhasil di create")
	}

	if check := db.Migrator().HasTable(&Users{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&Users{})
		fmt.Println("Table user berhasil di create")
	}
}
