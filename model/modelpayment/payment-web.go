package modelpayment

import (
	"github.com/haidityara/mtix/model/modelbook"
	"time"
)

type Request struct {
	Amount        string `json:"amount,omitempty"`
	PaymentMethod int    `json:"payment_method"`
	BookingID     uint   `json:"booking_id"`
}

type ResponseStore struct {
	ID            uint      `json:"id"`
	Amount        string    `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	BookingID     uint      `json:"booking_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ResponseGet struct {
	ID            uint                  `json:"id" gorm:"primaryKey"`
	Amount        string                `json:"amount"`
	PaymentMethod string                `json:"payment_method"`
	BookingID     uint                  `json:"booking_id"`
	Booking       modelbook.ResponseGet `json:"booking"`
	CreatedAt     time.Time             `json:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
}
