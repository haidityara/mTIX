package modelshow

import (
	"github.com/haidityara/mtix/model/modelcinemahall"
	"github.com/haidityara/mtix/model/modelmovie"
	"time"
)

type Request struct {
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	MovieID      uint   `json:"movie_id"`
	CinemaHallID uint   `json:"cinema_hall_id"`
	Price        string `json:"price"`
}

type ResponseStore struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	MovieID      uint      `json:"movie_id"`
	CinemaHallID uint      `json:"cinema_hall_id"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ResponseGet struct {
	ID           uint                        `json:"id" gorm:"primaryKey"`
	StartTime    string                      `json:"start_time"`
	EndTime      string                      `json:"end_time"`
	MovieID      uint                        `json:"movie_id"`
	Movie        modelmovie.Response         `json:"movie"`
	CinemaHallID uint                        `json:"cinema_hall_id"`
	CinemaHall   modelcinemahall.ResponseGet `json:"cinema_hall"`
	Price        string                      `json:"price"`
	CreatedAt    time.Time                   `json:"created_at"`
	UpdatedAt    time.Time                   `json:"updated_at"`
}
