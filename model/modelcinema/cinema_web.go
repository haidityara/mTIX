package modelcinema

import "github.com/haidityara/mtix/model/modelcinemacity"

type Request struct {
	Name   string `json:"name"`
	CityID uint   `json:"city_id"`
}

type ResponseStore struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ResponseGet struct {
	ID     uint                     `json:"id"`
	Name   string                   `json:"name"`
	CityID uint                     `json:"city_id"`
	City   modelcinemacity.Response `json:"cinema_city"`
}
