package entity

import "time"

type CinemaCity struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
