package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"size:36;not null;unique;primary_key;" json:"id"`
	Adreses	   	  []Address `gorm:"foreignKey:UserID"`
	FirstName     string `gorm:"size:255;not null;" json:"first_name"`
	LastName      string `gorm:"size:255;not null;" json:"last_name"`
	Email         string `gorm:"size:100;not null;unique;" json:"email"`
	Password      string `gorm:"size:100;not null;" json:"password"`
	RememberToken string `gorm:"size:255;not null" json:"remember_token"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt 	  gorm.DeletedAt `gorm:"index"`
}