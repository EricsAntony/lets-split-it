package user

import (
	"golang.org/x/crypto/bcrypt"
	"expense/pkg/graphModels"
)



func InsertUser(name, email, password string) (int, error) {
	// Encrypt the password.
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}
	stmt := ` insert into user (name , email , password) values (?,?,?)`
	row, Inserterr := database.Db.Exec(stmt, name, email, hashedpassword)
	if Inserterr != nil {
		return 0, Inserterr
	}

	userId, _ := row.LastInsertId()
	return int(userId), nil
}
