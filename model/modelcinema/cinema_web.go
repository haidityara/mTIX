package modelcinema

import (
	"github.com/haidityara/mtix/model/modelcinemacity"
	"time"
)

type Request struct {
	Name      string    `json:"name"`
	CityID    uint      `json:"city_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseStore struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGet struct {
	ID     uint                     `json:"id"`
	Name   string                   `json:"name"`
	CityID uint                     `json:"city_id"`
	City   modelcinemacity.Response `json:"cinema_city,omitempty"`
}
