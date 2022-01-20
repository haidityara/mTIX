package modelcinemaseat

import (
	"github.com/haidityara/mtix/model/modelcinemahall"
	"time"
)

type Request struct {
	SeatNum      int  `json:"seat_num"`
	CinemaHallID uint `json:"cinema_hall_id"`
}

type ResponseStore struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	SeatNum      int       `json:"seat_num"`
	CinemaHallID uint      `json:"cinema_hall_id"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type ResponseGet struct {
	ID           uint                         `json:"id" gorm:"primaryKey"`
	SeatNum      int                          `json:"seat_num"`
	CinemaHallID uint                         `json:"cinema_hall_id"`
	CinemaHall   *modelcinemahall.ResponseGet `json:"cinema_hall"`
	CreatedAt    time.Time                    `json:"created_at,omitempty"`
	UpdatedAt    time.Time                    `json:"updated_at,omitempty"`
}

type RequestAvailableSeat struct {
	CinemaHallID uint `json:"cinema_hall_id"`
	ShowID       uint `json:"show_id"`
}
