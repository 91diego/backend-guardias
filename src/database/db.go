package database

import (
	"fmt"

	"github.com/91diego/backend-guardias/config"
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
	// utils.EnvVariables()
	_, dbEnv := config.SetUp()

	dsn := dbEnv.DbUserName + ":" + dbEnv.DbPassword + "@tcp" + "(" + dbEnv.DbHost + ":" + dbEnv.DbPort + ")/" + dbEnv.DbName + "?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}
