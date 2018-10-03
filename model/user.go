package model

import (
	"time"
)

type User struct {
	ID int `json:"id" gorm:"primary_key"`
	Email     string     `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
	Password     string     `json:"password, omitempty" gorm:"not null; type:varchar(100)"`
	Role string `json:"role, omitempty" gorm:"not null; type:ENUM('admin', 'user', 'root')"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
}

func (User) TableName() string {
	return "users" // table name when succesfully migrate
}