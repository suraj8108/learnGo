package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "121aaa"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error present : %s", err)
	}

	db = result

}

func GetDB() *gorm.DB {
	return db
}
