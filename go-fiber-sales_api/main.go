package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"

	db "go-fiber-sales_api/config"
	routes "go-fiber/Routes"
)

func main() {
	fmt.Println("deve code app running...")

	db.Connect()

	app := fiber.New()
	app.Use(cors.New())
	//routing
	routes.Setup(app)
	app.Listen(":3030")

}
