package entity

import "time"

type Booking struct {
	ID        string    `json:"id"`
	Date      time.Time `json:"date"`
	Status    int       `json:"status"`
	UserID    string    `json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
