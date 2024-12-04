package main

import (
	"flag"
	"github.com/devkemc/fundamentos-golang/common"
	"github.com/devkemc/fundamentos-golang/emails"
	"github.com/devkemc/fundamentos-golang/orders"
	"github.com/devkemc/fundamentos-golang/payments"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	portFlag = "port"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "__ecommerce.db")
	if err != nil {
		log.Fatalln(err)
	}
	commonRepo := common.NewRepositorySqlx(db)

	paymentRepo := payments.NewPaymentRepositorySqlx(commonRepo)
	paymentService := payments.NewPaymentsServiceSimulator(paymentRepo)

	emailService := emails.NewEmailServiceSimulator()

	orderRepo := orders.NewOrderRepositorySqlx(commonRepo)
	orderServ := orders.NewOrderServiceV1(orderRepo, emailService, paymentService)
	orderHandler := orders.NewOrderHandler(orderServ)

	port := flag.String(portFlag, "8080", "port to server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")
	orders.SetupRoutes(apiV1, orderHandler)

	err = app.Listen(":" + *port)
	if err != nil {
		return
	}
}
