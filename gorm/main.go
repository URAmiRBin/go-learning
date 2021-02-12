package main

import (
	"gorm/handler"
	"gorm/model"
	"gorm/store"
	"log"

	echo "github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	if e == nil {
		log.Fatal()
	}

	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Migrator().AutoMigrate(&model.Student{}); err != nil {
		log.Fatal(err)
	}

	e.GET("/hello", handler.Hello)

	s := handler.NewStudent(store.NewSQLStudent(db))
	e.POST("/student", s.Create)
	e.GET("/student", s.Find)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}

}
