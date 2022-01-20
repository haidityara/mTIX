package entity

import "time"

type Payment struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Amount        string    `json:"amount"`
	PaymentMethod int       `json:"payment_method"`
	BookingID     uint      `json:"booking_id"`
	Booking       *Booking  `gorm:"foreignKey:BookingID"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
