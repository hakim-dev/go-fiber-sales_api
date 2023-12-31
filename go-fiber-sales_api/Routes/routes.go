package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber/controller"
)

func Setup(app *fiber.App) {
	//authentication routes
	app.Post("/cashiers/:cashierId/login", controller.Login)
	app.Post("/cashiers/:cashierId/logout", controller.Logout)
	app.Get("/cashiers/:cashierId/passcode", controller.Passcode)

	//Cashier routes
	app.Get("/cashiers", controller.CashiersList)
	app.Get("/cashiers/:cashierId", controller.GetCashierDetails)
	app.Post("/cashiers", controller.CreateCashier)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)
	app.Put("/cashiers/:cashierId", controller.UpdateCashier)

	//Category routes
	app.Get("/categories", controller.CategoryList)
	app.Get("/categories/:categoryId", controller.GetCategoryDetails)
	app.Post("/categories", controller.CreateCategory)
	app.Delete("/categories/:categoryId", controller.DeleteCategory)
	app.Put("/categories/:categoryId", controller.UpdateCategory)

	//Products routes
	app.Get("/products", controller.ProductList)
	app.Get("/products/:productId", controller.GetProductDetails)
	app.Post("/products", controller.CreateProduct)
	app.Delete("/products/:productId", controller.DeleteProduct)
	app.Put("/products/:productId", controller.UpdateProduct)

	//Payment routes
	app.Get("/payments", controller.PaymentList)
	app.Get("/payments/:paymentId", controller.GetPaymentDetails)
	app.Post("/payments", controller.CreatePayment)
	app.Delete("/payments/:paymentId", controller.DeletePayment)
	app.Put("/payments/:paymentId", controller.UpdatePayment)

	//Order routes
	app.Get("/orders", controller.OrdersList)
	app.Get("/orders/:orderId", controller.OrderDetail)
	app.Post("/orders", controller.CreateOrder)
	app.Post("/orders/subtotal", controller.SubTotalOrder)
	app.Get("/orders/:orderId/download", controller.DownloadOrder)
	app.Get("/orders/:orderId/check-download", controller.CheckOrder)

	//reports
	app.Get("/revenues", controller.GetRevenues)
	app.Get("/solds", controller.GetSolds)

}
