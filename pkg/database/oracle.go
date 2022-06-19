package database

import (
	"database/sql"
	"fmt"
	"log"
)

func NewOracleConnection(strConn string) (*sql.DB, error) {

	conn, err := sql.Open("oracle", strConn)
	if err != nil {
		return nil, fmt.Errorf("cannot open oracle connection, error: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database fail, error: %w", err)
	}

	log.Println("Oracle connection is activated")
	return conn, nil
}
