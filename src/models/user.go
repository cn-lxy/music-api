package models

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/cn-lxy/music-api/tools/db"
)

type User struct {
	Id       uint64
	NickName string
	Email    string
	Password string
}

// Insert a user into the database
func (u *User) Insert() (int64, error) {
	return db.Update("INSERT INTO users (nick_name, email, password) VALUES (?, ?, ?)", u.NickName, u.Email, u.Password)
}

// Update a user in the database
func (u *User) Update() (int64, error) {
	return db.Update("UPDATE users SET nick_name =?, email =?, password =? WHERE id =?", u.NickName, u.Email, u.Password, u.Id)
}

// Delete a user from the database
func (u *User) Delete() (int64, error) {
	// make sure the user exists
	if !u.Exists() {
		return 0, fmt.Errorf("user with id %v does not exist", u.Id)
	}
	return db.Update("DELETE FROM users WHERE id =?", u.Id)
}

// Check if a user exists in the database
func (u *User) Exists() bool {
	res, err := db.Query("SELECT id FROM users WHERE id = ?", u.Id)
	if err != nil {
		return false
	}
	return len(res) > 0
}

// Get a user from the database
func (u *User) Get() error {
	res, err := db.Query("SELECT id, nick_name, email, password FROM users WHERE id = ?", u.Id)
	if err != nil {
		return err
	}
	u.Id = res[0]["id"].(uint64)
	u.NickName = res[0]["nick_name"].(string)
	u.Email = res[0]["email"].(string)
	u.Password = res[0]["password"].(string)
	return nil
}

// GetByEmailOrNick By email and password or nick and password to get a user from the database
func (u *User) GetByEmailOrNick() error {
	var res []map[string]any
	var err error
	if u.Email != "" {
		res, err = db.Query("SELECT id, nick_name, email, password FROM users WHERE email = ? AND password = ?", u.Email, u.Password)
	} else if u.NickName != "" {
		res, err = db.Query("SELECT id, nick_name, email, password FROM users WHERE nick_name = ? AND password = ?", u.NickName, u.Password)
	}
	if err != nil {
		return err
	}
	if len(res) < 1 {
		return errors.New("User not found")
	}
	u.Id, err = strconv.ParseUint(res[0]["id"].(string), 10, 64)
	if err != nil {
		return err
	}
	u.NickName = res[0]["nick_name"].(string)
	u.Email = res[0]["email"].(string)
	u.Password = res[0]["password"].(string)
	return nil
}

// MapToUser convert map to user struct
func MapToUser(m map[string]interface{}) *User {
	// The map is a map of interface{} to interface{}. To get the value of an element, you must convert it to the appropriate type.
	return &User{
		Id:       m["id"].(uint64), // The map uses the key, "id", to get the value, which is then cast to a uint64.
		NickName: m["nick_name"].(string),
		Email:    m["email"].(string),
		Password: m["password"].(string),
	}
}
