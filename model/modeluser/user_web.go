package modeluser

import "time"

type Request struct {
	ID       uint   `json:"id,omitempty"`
	FullName string `json:"full_name" example:"jhondoe"`
	Email    string `json:"email" example:"test@example.com"`
	Phone    string `json:"phone" example:"+628123456789"`
	Password string `json:"password" example:"password"`
	Role     int    `json:"role" example:"1"`
}

type Response struct {
	ID        uint       `json:"id"  example:"1"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	CreatedAt *time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
	FullName  string     `json:"full_name" example:"jhondoe"`
	Email     string     `json:"email" example:"test@example.com"`
	Phone     string     `json:"phone" example:"+628123456789"`
	Role      int        `json:"role" example:"1"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password"`
}

type ResponseLogin struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJxd2Vxd2..."`
}
