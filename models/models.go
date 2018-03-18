package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Chat holds all the chat information in the chat room
type Chat struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" `
	DateCreated time.Time     `json:"dateCreated,omitempty"`
	UserName    string        `json:"username,omitempty"`
	Messages    []string      `json:"messages,omitempty"`
}

// User holds information of the user if he/she chooses to upgrade the account
type User struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" `
	FirstName      string        `json:"firstName,omitempty"`
	LastName       string        `json:"lastName,omitempty"`
	Description    string        `json:"description,omitempty"`
	AccountCreated time.Time     `json:"accountCreated,omitempty"`
}
