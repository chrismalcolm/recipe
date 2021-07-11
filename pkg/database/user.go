package database

import "github.com/chrismalcolm/recipe/pkg/model"

func (db *Database) AddUser(user model.User) (userID model.UserID, err error) {
	return userID, err
}

func (db *Database) EditUser(user model.User) (err error) {
	return err
}

func (db *Database) DeleteUser(userID model.UserID) (err error) {
	return err
}
