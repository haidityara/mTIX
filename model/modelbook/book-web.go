package modelbook

import "time"

type Request struct {
	Status        int  `json:"status"`
	UserID        uint `json:"user_id,omitempty"`
	BookingDetail []struct {
		CinemaSeatID uint `json:"cinema_seat_id"`
		ShowID       uint `json:"show_id"`
	} `json:"booking_detail"`
}

type ResponseStore struct {
	ID        uint      `json:"id"`
	Status    int       `json:"status"`
	UserID    uint      `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestUpdateStatus struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

type ResponseGet struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
	UserID uint `json:"user_id"`
	User   struct {
		ID        uint       `json:"id"  example:"1"`
		UpdatedAt *time.Time `json:"updated_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
		CreatedAt *time.Time `json:"created_at,omitempty" example:"2021-11-03T01:52:41.035Z"`
		FullName  string     `json:"full_name" example:"jhondoe"`
		Email     string     `json:"email" example:"test@example.com"`
		Phone     string     `json:"phone" example:"+628123456789"`
	} `json:"user"`
	BookingDetail []struct {
		ID           uint      `json:"id`
		NumberOFSeat int       `json:"number_of_seat"`
		Price        string    `json:"price"`
		CinemaSeatID uint      `json:"cinema_seat_id"`
		ShowID       uint      `json:"show_id"`
		BookingID    uint      `json:"booking_id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	} `json:"booking_detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
