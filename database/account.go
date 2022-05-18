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
	query := "SELECT balance FROM accounts WHERE id = $1"
	row := db.QueryRow(query, accountId)
	var balance int
	err := row.Scan(&balance)
	return balance, err
}

func DepositAmount(db *sql.DB, accountId uuid.UUID, amount int) error {
	// todo
	return nil
}
