package migrator

import (
	"fmt"
	"log"
	"xapiens-project-gin-graphql/models"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {

	// drop table user
	db.Migrator().DropTable(&models.Vendor{})
	if check := db.Migrator().HasTable(&models.Vendor{}); !check {
		db.Migrator().CreateTable(&models.Vendor{})
		fmt.Println("Table vendor berhasil di create")
	} else {
		log.Println(check)
	}

	db.Migrator().DropTable(&models.Employee{})
	if check := db.Migrator().HasTable(&models.Employee{}); !check {
		db.Migrator().CreateTable(&models.Employee{})
		fmt.Println("Table employee berhasil di create")
	}

	/* if check := db.Migrator().HasTable(&Users{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&Users{})
		fmt.Println("Table user berhasil di create")
	} */
}
