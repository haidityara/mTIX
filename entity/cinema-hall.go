package entity

import "time"

type CinemaHall struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	TotalSeat  int     `json:"total_seat"`
	CinemaID   uint    `json:"cinema_id"`
	Cinema     *Cinema `gorm:"foreignKey:CinemaID"`
	CinemaSeat []CinemaSeat
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
