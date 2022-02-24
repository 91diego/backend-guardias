package database

import (
	"fmt"

	"github.com/91diego/backend-guardias/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	utils.EnvVariables()

	dbUserName := "calendar_user"        //os.Getenv("DB_USERNAME")
	dbPassword := " id3xGu@r1asAs3s0r3s" //os.Getenv("DB_PASSWORD")
	dbHost := "localhost"                // os.Getenv("DB_HOST")
	dbPort := "3306"                     //os.Getenv("DB_PORT")
	dbName := "guardias"                 //os.Getenv("DB_NAME")

	dsn := dbUserName + ":" + dbPassword + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}
