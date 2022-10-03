package main

import (
	"payroll/internal/config"
	"payroll/internal/config/db"
	"payroll/internal/controller"
	"payroll/internal/model"
	"payroll/internal/ms"
	"payroll/internal/storage"
	"payroll/internal/usecase"

	// "os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	err := godotenv.Load("./etc/local.env")
	if err != nil { 
		log.Printf("please consider environment variables: %s", err)
	}

	msService := newMSService(echo.New())
	msService.InitRoutes()
	msService.Server()

}

func newMSService(e *echo.Echo) *ms.Service {

	config, err := config.New()
	if err != nil {
		panic(err)
	}

	conn := ("host=localhost user=postgres password=1234 dbname=payroll port=8000 sslmode=disable TimeZone=Asia/Shanghai")

	db, err := db.New(conn)
	if err != nil {
		log.Fatalf("cannot connect to DB: %w", err)


	}

	// Create table for `User`
	db.Migrator().CreateTable(&model.Payroll{}, &model.Employee{}, &model.JobDepartment{}, &model.Leave{}, &model.Salary{})

	db.AutoMigrate(&model.Payroll{}, &model.Employee{}, &model.JobDepartment{}, &model.Leave{}, &model.Salary{})

	pg := storage.NewService(db)

	usecases := usecase.NewService(config, pg)

	return &ms.Service{
		Echo:   e,
		HTTP:   controller.NewService(usecases),
		Config: config,
	}

}
