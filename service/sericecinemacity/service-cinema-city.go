package sericecinemacity

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelcinemacity"
	"github.com/haidityara/mtix/repository/repositorycinemacity"
	"github.com/haidityara/mtix/validation"
	"github.com/jinzhu/copier"
)

type ServiceCinemaCity interface {
	Create(request modelcinemacity.Request) (modelcinemacity.Response, error)
	GetByID(id string) (modelcinemacity.Response, error)
}

type service struct {
	repo repositorycinemacity.RepositoryCinemaCity
}

func (s service) Create(request modelcinemacity.Request) (modelcinemacity.Response, error) {
	// validation
	err := validation.ValidateStoreCinemaCity(request)
	if err != nil {
		return modelcinemacity.Response{}, err
	}

	cinemaCity := entity.CinemaCity{}
	copier.Copy(&cinemaCity, &request)
	// store to repo
	create, err := s.repo.Create(cinemaCity)
	if err != nil {
		return modelcinemacity.Response{}, err
	}

	resp := modelcinemacity.Response{}
	copier.Copy(&resp, &create)
	return resp, nil
}

func (s service) GetByID(id string) (modelcinemacity.Response, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return modelcinemacity.Response{}, err
	}
	resp := modelcinemacity.Response{}
	copier.Copy(&resp, &data)
	return resp, nil
}

func New(repo repositorycinemacity.RepositoryCinemaCity) ServiceCinemaCity {
	return &service{repo: repo}
}
