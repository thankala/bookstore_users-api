module github.com/thankala/bookstore_users-api

go 1.15

//Ive seen stupid things in my life but this is too much. Future me dont forget this crap
replace github.com/thankala/bookstore_auth-go => /home/thankala/go/src/github.com/thankala/bookstore_auth-go
require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofiber/fiber/v2 v2.2.3
	github.com/joho/godotenv v1.3.0
	github.com/thankala/bookstore_auth-go v0.0.0-20201206201607-1edbbcfb38c5
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20201204225414-ed752295db88 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)
