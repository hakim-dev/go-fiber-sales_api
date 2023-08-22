package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	routes "go-fiber/Routes"
	db "go-fiber/config"
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
