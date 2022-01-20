package modelcinemacity

import "time"

type Request struct {
	Name      string    `json:"name"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Response struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	Zip       string    `json:"zip"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
