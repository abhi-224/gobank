package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccountById(int) error
	GetAccount() ([]*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=postgres dbname=gobank_db password=changeinprod sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createTable("accounts")
}

func (s *PostgresStore) createTable(name string) error {
	ddl := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number INT NOT NULL UNIQUE,
		balance NUMERIC(12, 2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`, name)

	_, err := s.db.Exec(ddl)
	return err
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	return s.db.QueryRow(`
	INSERT INTO accounts(first_name, last_name, number, balance)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`, a.FirstName, a.LastName, a.Number, a.Balance).Scan(&a.Id)

}

func (s *PostgresStore) GetAccountById(id int) error {
	return nil
}

func (s *PostgresStore) GetAccount() ([]*Account, error) {
	rows, err := s.db.Query(`
	SELECT id, first_name, last_name, number, balance
	FROM accounts
	ORDER BY id;
	`)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		err := rows.Scan(
			&account.Id,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil

}

func (s *PostgresStore) UpdateAccount(a *Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
