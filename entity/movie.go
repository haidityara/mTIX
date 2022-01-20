package entity

import "time"

type Movie struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Language    string        `json:"language"`
	RealiseDate time.Time     `json:"realise_date"`
	Country     string        `json:"country"`
	Genres      string        `json:"genres"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
