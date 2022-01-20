package entity

import "time"

type CinemaHall struct {
	ID         string  `json:"id"`
	TotalSeat  int     `json:"total_seat"`
	CinemaID   string  `json:"cinema_id"`
	Cinema     *Cinema `gorm:"foreignKey:CinemaID"`
	CinemaSeat []CinemaSeat
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}