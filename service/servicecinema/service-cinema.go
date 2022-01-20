package servicecinema

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelcinema"
	"github.com/haidityara/mtix/repository/repositorycinema"
	"github.com/haidityara/mtix/validation"
	"github.com/jinzhu/copier"
)

type ServiceCinema interface {
	Create(request modelcinema.Request) (modelcinema.ResponseStore, error)
	GetByID(id string) (modelcinema.ResponseGet, error)
}

type service struct {
	repo repositorycinema.RepositoryCinema
}

func (s *service) Create(request modelcinema.Request) (modelcinema.ResponseStore, error) {
	err := validation.ValidateStoreCinema(request)
	if err != nil {
		return modelcinema.ResponseStore{}, err
	}
	cinema := entity.Cinema{}
	copier.Copy(&cinema, &request)
	create, err := s.repo.Create(cinema)
	if err != nil {
		return modelcinema.ResponseStore{}, err
	}
	resp := modelcinema.ResponseStore{}
	copier.Copy(&resp, &create)
	return resp, nil
}

func (s *service) GetByID(id string) (modelcinema.ResponseGet, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return modelcinema.ResponseGet{}, err
	}
	resp := modelcinema.ResponseGet{}
	copier.Copy(&resp, &data)
	return resp, nil

}

func New(repo repositorycinema.RepositoryCinema) ServiceCinema {
	return &service{repo: repo}
}
