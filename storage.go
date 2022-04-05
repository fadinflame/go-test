package main

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (storage *SQLiteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS customers(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        balance FLOAT NOT NULL
    );
    `

	_, err := storage.db.Exec(query)
	return err
}

func (storage *SQLiteRepository) CreateCustomer(customer Customer) (*Customer, error) {
	var query = `INSERT INTO customers (name, balance) VALUES (?, 0)`
	var res, err = storage.db.Exec(query, customer.Name)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	customer.ID = id

	return &customer, nil
}

func (storage *SQLiteRepository) getCustomerByName(name string) (*Customer, error) {
	var query = `SELECT * FROM customers WHERE name = ?;`
	var res = storage.db.QueryRow(query, name)
	var customer Customer
	if err := res.Scan(&customer.ID, &customer.Name, &customer.balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &customer, nil

}


func (storage *SQLiteRepository) depositToBalance (name string, amount float64) (bool, error) {
	var query = `UPDATE customers SET balance = balance + ? WHERE name = ?`
	res, err := storage.db.Exec(query, amount, name)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, ErrUpdateFailed
	}

	return true, nil
}