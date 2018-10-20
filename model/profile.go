package model

import (
	"time"
)

//go:generate goqueryset -in profile.go

// Profile struct represent profile model. Next line (gen:qs) is needed to autogenerate ProfileQuerySet.
// gen:qs
type Profile struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	UserID    int64      `json:"user_id,omitempty" gorm:"type:bigint REFERENCES users(id)"`
	Address   string     `json:"address,omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string     `json:"lastname,omitempty" gorm:"not null; type:varchar(100)"`
	FirstName string     `json:"firstname,omitempty" gorm:"not null; type:varchar(100)"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
	User      User       `gorm:"foreignkey:UserRefer"`
}

func (Profile) TableName() string {
	return "profiles" // table name when succesfully migrate
}
