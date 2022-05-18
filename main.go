package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/MonduCareers/go-sample-app/handlers"
)

func main() {
	db, err := sql.Open("postgres", "postgres://root:mondu123@localhost/mondu_dev?sslmode=disable")
	if err != nil {
		fmt.Println("cannot initialize db", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("can't connect to DB", err)
		os.Exit(1)
	}

	http.Handle("/account/create", handlers.NewAccountCreateHandler(db))
	http.Handle("/account/balance", handlers.NewAccountBalanceHandler(db))
	http.Handle("/account/deposit", handlers.NewAccountDepositHandler(db))

	fmt.Println("Starting the server on 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed to listen on 8080", err)
		os.Exit(1)
	}
}
