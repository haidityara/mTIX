package entity

import "time"

type BookingDetail struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	NumberOFSeat int         `json:"number_of_seat"`
	Price        string      `json:"price"`
	CinemaSeatID string      `json:"cinema_seat_id"`
	CinemaSeat   *CinemaSeat `gorm:"foreignKey:CinemaSeatID"`
	ShowID       string      `json:"show_id"`
	Show         *Show       `gorm:"foreignKey:ShowID"`
	BookingID    string      `json:"booking_id"`
	Booking      *Booking    `gorm:"foreignKey:BookingID"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
