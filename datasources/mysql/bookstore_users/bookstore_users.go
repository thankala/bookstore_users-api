package bookstore_users

import (
	"database/sql"
	"fmt"
	"github.com/thankala/bookstore_users-api/domain/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	usersDB *sql.DB
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user",
			"password",
			"127.0.0.1",
			"3307",
			"bookstore_users",
	)
	usersDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database successfully configures")

	usersDB.AutoMigrate(&users.User{})
}
