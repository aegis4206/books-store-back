package model

import (
	"books-store/utils"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func CheckEmailAndPassword(email string, password string) (*User, error) {
	sqlStr := "select id,email from users where email=$1 and password=$2"
	row := utils.Db.QueryRow(sqlStr, email, password)
	user := &User{}
	row.Scan(&user.Id, &user.Email)
	return user, nil
}

func CheckEmail(email string) (*User, error) {
	sqlStr := "select id,email,password from users where email=$1 "
	row := utils.Db.QueryRow(sqlStr, email)
	user := &User{}
	row.Scan(&user.Id, &user.Email, &user.Password)
	return user, nil
}

func SaveUser(email string, password string) error {
	sqlStr := "insert into users(email,password) values($1,$2)"
	_, err := utils.Db.Exec(sqlStr, email, password)
	if err != nil {
		return err
	}
	return nil
}
