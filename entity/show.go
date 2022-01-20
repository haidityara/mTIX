package entity

import "time"

type Show struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	StartTime    time.Time   `json:"start_time"`
	EndTime      time.Time   `json:"end_time"`
	DateShow     time.Time   `json:"date"`
	MovieID      uint        `json:"movie_id"`
	Movie        *Movie      `gorm:"foreignKey:MovieID"`
	CinemaHallID uint        `json:"cinema_hall_id"`
	CinemaHall   *CinemaHall `gorm:"foreignKey:CinemaHallID"`
	Price        string      `json:"price"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
