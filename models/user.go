package models

import "time"

type User struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,string,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}
