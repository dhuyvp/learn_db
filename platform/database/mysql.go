package database

import (
	"fmt"
	"learn_db/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Variable struct {
	Db *sqlx.DB
	Tx *sqlx.Tx
}

func New() (*Variable, error) {

	db, err := MySQLConnection()

	if err != nil {
		return nil, err
	}

	return &Variable{
		Db: db,
		Tx: nil,
	}, nil
}

func MySQLConnection() (*sqlx.DB, error) {
	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("mysql", mysqlConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}
	return db, nil
}
