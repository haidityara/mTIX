package entity

import "time"

type Booking struct {
	ID            uint  `json:"id" gorm:"primaryKey"`
	Status        int   `json:"status"`
	UserID        uint  `json:"user_id"`
	User          *User `gorm:"foreignKey:UserID"`
	BookingDetail []BookingDetail
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
