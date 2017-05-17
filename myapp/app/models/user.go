package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Tell     string `json:"tell"`
	Sex      string `json:"sex"`
	EventID  uint   `json:"event_id" gorm:"not null"`
}
