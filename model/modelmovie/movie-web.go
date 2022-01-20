package modelmovie

import "time"

type Request struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Language    string        `json:"language"`
	RealiseDate string        `json:"realise_date"`
	Country     string        `json:"country"`
	Genres      string        `json:"genres"`
}

type Response struct {
	ID          uint          `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	Language    string        `json:"language"`
	RealiseDate time.Time     `json:"realise_date"`
	Country     string        `json:"country"`
	Genres      string        `json:"genres"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
}
