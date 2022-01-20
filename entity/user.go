package entity

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Phone     string    `json:"phone"`
	Role      int       `json:"role" gorm:"not null,default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
