package main

import (
	"log"
)

func main() {
	dbCfgs := Postgres{
		Host:     "localhost",
		Name:     "postgres",
		User:     "postgres",
		Password: "123",
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

	stmt := "DELETE FROM users where id=1"
	_, err = connDB.Query(stmt)
	if err != nil {
		log.Printf("error deleting user in db: %v", err.Error())
		return
	}
}
