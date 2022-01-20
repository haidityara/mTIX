package entity

import "time"

type CinemaSeat struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	SeatNum      int         `json:"seat_num"`
	CinemaHallID uint        `json:"cinema_hall_id"`
	CinemaHall   *CinemaHall `gorm:"foreignKey:CinemaHallID"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}
