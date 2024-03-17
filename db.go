package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Host     string `json:"Host"`
	Name     string `json:"Name"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Schema   string `json:"Schema"`
}

var dbConn *sql.DB
var dbConfig *Postgres

func InitDB(postgres Postgres) (*sql.DB, error) {
	if dbConn == nil {
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", postgres.Host, "IGOR", postgres.Password, postgres.Name)
		d, err := sql.Open("postgres", conn)
		if err != nil {
			return nil, err
		}

		if postgres.Schema != "" {
			_, err = d.Exec(fmt.Sprintf(`CREATE SCHEMA IF NOT EXISTS "%v"`, postgres.Schema))
			if err != nil {
				return nil, err
			}
		}

		dbConn = d
		dbConfig = &postgres
	}
	return dbConn, nil
}
