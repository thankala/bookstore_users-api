package bookstore_users

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	gorm.Model
	Id        int64  `gorm:"primaryKey" json:"Id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Status    string `json:"status"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

const (
	mysqlUsersUsername = "MYSQL_USERNAME"
	mysqlUsersPassword = "MYSQL_PASSWORD"
	mysqlUsersHost     = "MYSQL_HOST"
	mysqlUsersPort     = "MYSQL_PORT"
	mysqlUsersSchema   = "MYSQL_SCHEMA"
)

var (
	Client   *gorm.DB
	_        = godotenv.Load("variables.env")
	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	port     = os.Getenv(mysqlUsersPort)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC&time_zone=UTC",
		username,
		password,
		host,
		port,
		schema,
	)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	Client = db.Session(&gorm.Session{PrepareStmt: true})
	log.Println("Database successfully configured")
	err = Client.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to auto migrate")
	}
}
