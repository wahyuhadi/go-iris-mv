package model

import (
	"time"
)

type Profile struct {
	//gorm.Model
	ID        int        `json:"id" gorm:"primary_key"`
	UserID    int        `json:"user_id, omitempty" gorm:"not null"`
	Address   string     `json:"address, omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string     `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	FirstName string     `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
	User      User       `gorm:"foreignkey:UserRefer"`
}

func (Profile) TableName() string {
	return "profiles" // table name when succesfully migrate
}
