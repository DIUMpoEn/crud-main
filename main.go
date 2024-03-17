package main

import (
	"fmt"
	"log"
)

type accaunt struct {
	id       int
	login    string
	password string
	email    string
}

func main() {
	dbCfgs := Postgres{
		Host:     "localhost",
		Name:     "postgres",
		User:     "postgres",
		Password: "12345678",
	}

	/// crud
	/// c - create
	/// r - read
	/// u - update
	/// d - delete

	connDB, err := InitDB(dbCfgs)
	if err != nil {
		log.Printf("error initializing to db: %v", err.Error())
		return
	}

	//stmt := "INSERT INTO users(login, password, email) VALUES ('igor', '12345', 'chapligin@mail.ru') RETURNING login, password, email"
	//_, err = connDB.Query(stmt)
	//if err != nil {
	//	log.Printf("error creating user in db: %v", err.Error())
	//	return
	//}

	//stmt := "DELETE FROM users where id=1"
	//_, err = connDB.Query(stmt)
	//if err != nil {
	//	log.Printf("error deleting user in db: %v", err.Error())
	//	return

	stmt := "SELECT * FROM users"
	rows, err := connDB.Query(stmt)
	if err != nil {
		log.Printf("error selectting user in db: %v", err.Error())
		return
	}
	date := []accaunt{}

	for rows.Next() {
		p := accaunt{}
		err := rows.Scan(&p.id, &p.login, &p.password, &p.email)
		if err != nil {
			fmt.Println(err)
			continue
		}
		date = append(date, p)
	}
	for _, p := range date {
		fmt.Println(p.id, p.login, p.password, p.email)
	}

}
