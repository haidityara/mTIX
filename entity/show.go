package entity

import "time"

type Show struct {
	ID           string      `json:"id"`
	StartDate    time.Time   `json:"start_date"`
	EndDate      time.Time   `json:"end_date"`
	MovieID      string      `json:"movie_id"`
	Movie        *Movie      `gorm:"foreignKey:MovieID"`
	CinemaHallID string      `json:"cinema_hall_id"`
	CinemaHall   *CinemaHall `gorm:"foreignKey:CinemaHallID"`
	Price        string      `json:"price"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
