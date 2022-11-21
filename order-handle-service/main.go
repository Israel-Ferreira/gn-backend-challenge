package main

import (
	"fmt"
    "github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/internal/config"
    "log"
	"net/http"

	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/internal/handler"
	"github.com/Israel-Ferreira/gn-backend-challenge/order-handle-service/pkg/middlewares"
)

func init() {
    config.LoadEnv()
    config.ConfigRabbitMq()
}

func main() {
    config.DbConfiguration()

	port := fmt.Sprintf(":%d", 9000)

    defer config.ConnRabbitMq.Close()

	http.HandleFunc("/orders", middlewares.JsonMiddleware(handler.OrderHandler))

	log.Printf("Servidor Iniciado na porta %s \n", port)

	log.Fatalln(http.ListenAndServe(port, nil))
}
