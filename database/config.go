package database

import (
	"fmt"
	"mygram/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// config db
const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "db-mygram"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {

	// conn db
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", username, password, hostname, dbname))
	if err != nil {
		panic(err)
	}

	// auto migrate models to database table
	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.SocialMed{})
	db.Debug().AutoMigrate(models.Photo{})
	db.Debug().AutoMigrate(models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
