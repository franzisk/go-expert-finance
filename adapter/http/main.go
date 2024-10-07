package http

import (
	"net/http"

	"github.com/franzisk/go-expert-finance/adapter/http/actuator"
	"github.com/franzisk/go-expert-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init serve para iniciar o serviço REST da aplicação
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)

	http.HandleFunc("/transactions/create", transaction.CraeteTransaction)

	http.HandleFunc("/health", actuator.HealthCheck)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
