package models

import "time"

type Address struct {
	ID         string `gorm:"size:36;not null;unique;primary_key;" json:"id"`
	User 	   User
	UserID     string `gorm:"size:36;index", json:"user_id"` 
	Name       string `gorm:"size:255" json:"name"`
	IsPrimary  bool
	CityID     string `gorm:"size:100", json:"city_id"`
	ProvinceID string `gorm:"size:100", json:"province_id"`
	Address1   string `gorm:"size:255", json:"address_1"`
	Address2   string `gorm:"size:255", json:"address_2"`
	Phone      string `gorm:"size:20", json:"phone"`
	Email      string `gorm:"size:100", json:"email"`
	PostalCode string `gorm:"size:10", json:"postal_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}