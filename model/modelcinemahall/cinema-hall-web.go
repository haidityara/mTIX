package modelcinemahall

import (
	"github.com/haidityara/mtix/model/modelcinema"
	"time"
)

type Request struct {
	Name      string `json:"name"`
	TotalSeat int    `json:"total_seat"`
	CinemaID  uint   `json:"cinema_id"`
}

type ResponseStore struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	TotalSeat int       `json:"total_seat"`
	CinemaID  uint      `json:"cinema_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGet struct {
	ID         uint                    `json:"id" gorm:"primaryKey"`
	Name       string                  `json:"name"`
	TotalSeat  int                     `json:"total_seat"`
	CinemaID   uint                    `json:"cinema_id"`
	CreatedAt  time.Time               `json:"created_at"`
	UpdatedAt  time.Time               `json:"updated_at"`
	Cinema     modelcinema.ResponseGet `json:"cinema"`
	CinemaSeat []struct {
		ID        uint      `json:"id" gorm:"primaryKey"`
		SeatNum   int       `json:"seat_num"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	} `json:"cinema_seat,omitempty"`
}
