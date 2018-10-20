package model

import (
	"time"
)

//go:generate goqueryset -in user.go

// User struct represent user model. Next line (gen:qs) is needed to autogenerate UserQuerySet.
// gen:qs
type User struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	Role      string     `json:"role,omitempty" gorm:"not null"`
	Email     string     `json:"email" gorm:"not null; size:255"`
	Password  string     `json:"password" gorm:"not null; size:255"` // Default size for string is 255, reset it with this tag
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" sql:"index"`
	//Profile   *Profile   `json:"profile"` //`gorm:"foreignkey:UserID;association_foreignkey:Refer"` // One-To-Many relationship (has many - use Email's UserID as foreign key)
}

func (User) TableName() string {
	return "users" // table name when succesfully migrate
}
