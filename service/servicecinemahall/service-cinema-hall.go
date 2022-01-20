package servicecinemahall

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelcinemahall"
	"github.com/haidityara/mtix/repository/repositorycinemahall"
	"github.com/jinzhu/copier"
)

type ServiceCinemaHall interface {
	Create(request modelcinemahall.Request) (modelcinemahall.ResponseStore, error)
	GetByID(id string) (modelcinemahall.ResponseGet, error)
}

type service struct {
	repo repositorycinemahall.RepositoryCinemaHall
}

func (s *service) Create(request modelcinemahall.Request) (modelcinemahall.ResponseStore, error) {
	cinemaHall := entity.CinemaHall{}
	copier.Copy(&cinemaHall, &request)
	response, err := s.repo.Create(cinemaHall)
	if err != nil {
		return modelcinemahall.ResponseStore{}, err
	}
	resp := modelcinemahall.ResponseStore{}
	copier.Copy(&resp, &response)
	return resp, nil
}

func (s *service) GetByID(id string) (modelcinemahall.ResponseGet, error) {
	response, err := s.repo.GetByID(id)
	if err != nil {
		return modelcinemahall.ResponseGet{}, err
	}
	resp := modelcinemahall.ResponseGet{}
	copier.Copy(&resp, &response)
	return resp, nil
}

func New(repo repositorycinemahall.RepositoryCinemaHall) ServiceCinemaHall {
	return &service{
		repo: repo,
	}
}
