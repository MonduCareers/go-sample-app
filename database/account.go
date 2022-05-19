package database

import (
	"database/sql"

	"github.com/google/uuid"
)

func CreateAccount(db *sql.DB) (uuid.UUID, error) {
	row := db.QueryRow("INSERT INTO accounts DEFAULT VALUES RETURNING id")
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

func GetAccountBalance(db *sql.DB, accountId uuid.UUID) (int, error) {
	return getBalance(db, accountId)
}

func DepositAmount(db *sql.DB, accountId uuid.UUID, amount int) error {
	existingBalance, err := getBalance(db, accountId)
	if err != nil {
		return err
	}
	newBalance := existingBalance + amount
	query := "UPDATE accounts SET balance = $1 WHERE id = $2"
	_, err = db.Exec(query, newBalance, accountId)
	return err
}

func getBalance(db *sql.DB, accountId uuid.UUID) (int, error) {
	query := "SELECT balance FROM accounts WHERE id = $1"
	row := db.QueryRow(query, accountId)
	var balance int
	err := row.Scan(&balance)
	return balance, err
}
