package entity

import "time"

type CinemaSeat struct {
	ID           string      `json:"id"`
	SeatNum      int         `json:"seat_num"`
	CinemaHallID string      `json:"cinema_hall_id"`
	CinemaHall   *CinemaHall `gorm:"foreignKey:CinemaHallID"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
