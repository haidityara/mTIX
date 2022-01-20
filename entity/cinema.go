package entity

import "time"

type Cinema struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name"`
	CityID    string      `json:"city_id"`
	City      *CinemaCity `gorm:"foreignKey:CityID"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
