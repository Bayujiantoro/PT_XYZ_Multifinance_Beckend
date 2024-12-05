package main

import (
	"fmt"
	"log"
	"os"
	"pt-xyz-multifinance/connection"
	"pt-xyz-multifinance/migration"
	"pt-xyz-multifinance/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))


	connection.DatabaseConnection()
	migration.RunMigration()

	routes.RouterInit(e.Group("/api/v1"))

	fmt.Println("Runing on Port : ", port)
	e.Logger.Fatal(e.Start("localhost:"+ port))
}