package model

import (
	"time"
)

type User struct {
	ID int `json:"id" gorm:"primary_key"`

	FirstName string     `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string     `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	Email     string     `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
}

func (User) TableName() string {
	return "users" // table name when succesfully migrate
}