package users

import (
	"database/sql"
	"log"
)

type Queries struct {
	GetDataUserByID *sql.Stmt
	ValidateUser    *sql.Stmt
}

func prepare(query string, db *sql.DB) *sql.Stmt {
	s, err := db.Prepare(query)
	if err != nil {
		log.Println("failed to prepare query", query, err)
	}
	return s
}

func NewQueries(dbMaster, dbSlave *sql.DB) *Queries {
	q := &Queries{
		GetDataUserByID: prepare(getDataUserByID, dbSlave),
		ValidateUser:    prepare(validateUser, dbSlave),
	}
	return q
}

const (
	getDataUserByID = `
	SELECT
	user_id
	, name
	, password
	, last_login
	, birth_date
	, address
	, gender
	role_id
	FROM users
	WHERE user_id = ?;`
	validateUser = `
	SELECT user_id, username from users where username = ? and password = ?;`
)
