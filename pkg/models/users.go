package models

import (
	"errors"
)

var (
	selectUserByUsername = `SELECT * FROM users WHERE username=$1;`
	selectUserByID       = `SELECT * FROM users WHERE id=$1;`
	selectUsers          = `SELECT * FROM users;`
	insertUser           = `INSERT INTO users (username, password, role_id) VALUES (:username, :password, :roleid) RETURNING id;`
	deleteUserByID       = `DELETE FROM users WHERE id=$1`
	updateUser           = `UPDATE users SET username=:username, password=:password WHERE id=:id `
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   string `json:"role_id" db:"role_id"`
}

// Users is a structure representing a set of users. Its purpose is mainly to
// ease the JSON serialization (to return a root key).
type Users struct {
	Users []User `json:"users"`
}

// CreateNewUser creates a new user, persists it and returns it.
func (r DatabaseRepository) GetUserByUsername(username string) (*User, error) {
	u := &User{}
	err := r.db.Get(u, selectUserByUsername, username)
	return u, err
}

func (r DatabaseRepository) GetUserByID(id int) (*User, error) {
	u := &User{}
	err := r.db.Get(u, selectUserByID, id)
	return u, err
}

func (r DatabaseRepository) DeleteUserByID(id int) error {
	_, err := r.db.Exec(deleteUserByID, id)
	return err
}

// CreateUser blah
func (r DatabaseRepository) CreateUser(user User) (User, error) {
	var id int
	rows, err := r.db.NamedQuery(insertUser, user)
	if err != nil {
		return user, err
	}
	if rows.Next() {
		rows.Scan(&id)
		user.ID = id
		return user, nil
	}
	return user, errors.New("could not find new id")
}

// GetUsers
func (r DatabaseRepository) GetUsers() (Users, error) {
	users := &[]User{}
	err := r.db.Select(users, selectUsers)
	return Users{*users}, err
}
