// @title RESTful API Contact
// @version 1.0
// @description Ini adalah dokumentasi API kontak menggunakan Golang, Fiber, dan Swagger.
// @host localhost:3000
// @BasePath /api
package main

import (
	"crud/config"
	"crud/database"
	"crud/routes"

	_ "crud/docs" // ganti dengan nama module kamu

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger" // yang benar
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
