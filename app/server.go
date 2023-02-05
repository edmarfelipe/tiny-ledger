package main

import (
	"log"

	"github.com/edmarfelipe/tiny-ledger/db"
	"github.com/edmarfelipe/tiny-ledger/usecase/account/block"
	"github.com/edmarfelipe/tiny-ledger/usecase/account/create"
	"github.com/edmarfelipe/tiny-ledger/usecase/account/getbalance"
	createPerson "github.com/edmarfelipe/tiny-ledger/usecase/person/create"
	createTransaction "github.com/edmarfelipe/tiny-ledger/usecase/transaction/create"
	"github.com/edmarfelipe/tiny-ledger/usecase/transaction/gettransactions"

	"github.com/dlmiddlecote/sqlstats"
	"github.com/edmarfelipe/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer() chi.Router {
	conn, err := db.Connect()
	if err != nil {
		log.Println("fail to connect to database", err)
	}

	prometheus.MustRegister(
		sqlstats.NewStatsCollector("postgres", conn),
	)

	accountDB := db.NewAccountDB(conn)
	transactionDB := db.NewTransactionDB(conn)
	personDB := db.NewPersonDB(conn)

	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(chiprometheus.NewMiddleware("api"))

	r.Handle("/metrics", promhttp.Handler())
	r.Post("/persons", handle(createPerson.New(personDB)))
	r.Post("/accounts", handle(create.New(accountDB, personDB)))

	// /block
	// /unblock
	r.Patch("/accounts/{account-id}", handle(block.New(accountDB)))

	// /withdraw
	r.Post("/accounts/{account-id}/deposit", handle(createTransaction.New(transactionDB)))
	r.Get("/accounts/{account-id}/balance", handle(getbalance.New(accountDB)))
	r.Get("/accounts/{account-id}/transactions", handle(gettransactions.New(transactionDB)))

	return r
}
