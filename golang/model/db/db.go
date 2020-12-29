package db_mod

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"unique_index"`
	Password	string
	Email		string
}

type Question struct {
	gorm.Model
	UserID		uint
	Data		string
}

type Answer struct {
	gorm.Model
	Question	uint
	Data		string
}