package db

import (
	"crud/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", postgres.Host, postgres.Name, postgres.Password, postgres.Name)
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

func GetUserFromDB(userId string) []models.Account {
	dbCfgs := Postgres{
		Host:     "localhost",
		Name:     "postgres",
		User:     "postgres",
		Password: "12345678",
	}

	connDB, err := InitDB(dbCfgs)
	if err != nil {
		log.Printf("error initializing to db: %v", err.Error())
		return nil
	}

	stmt := "SELECT * FROM users WHERE id = $1"
	rows, err := connDB.Query(stmt, userId)
	if err != nil {
		log.Printf("error selectting user in db: %v", err.Error())
		return nil
	}
	data := []models.Account{}

	for rows.Next() {
		p := models.Account{}
		err := rows.Scan(&p.Id, &p.Login, &p.Password, &p.Email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data = append(data, p)
	}
	for _, p := range data {
		fmt.Println(p.Id, p.Login, p.Password, p.Email)
	}

	return data
}

func DropUserFromDB(userId string) error {
	dbCfgs := Postgres{
		Host:     "localhost",
		Name:     "postgres",
		User:     "postgres",
		Password: "12345678",
	}

	connDB, err := InitDB(dbCfgs)
	if err != nil {
		log.Printf("error initializing to db: %v", err.Error())
		return nil
	}

	stmt := "delete  FROM users WHERE id = $1"
	_, err = connDB.Query(stmt, userId)
	if err != nil {
		log.Printf("error deleting user in db: %v", err.Error())
		return err
	}
	return nil

}
