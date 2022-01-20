package sericecinemacity

import (
	"github.com/haidityara/mtix/model/modelcinemacity"
	"github.com/haidityara/mtix/repository/repositorycinemacity"
)

type ServiceCinemaCity interface {
	Create(request modelcinemacity.Request) (modelcinemacity.Response, error)
	GetByID(id string) (modelcinemacity.Response, error)
}

type service struct {
	repo repositorycinemacity.RepositoryCinemaCity
}

func (s service) Create(request modelcinemacity.Request) (modelcinemacity.Response, error) {
	cinemaCity := enti
}

func (s service) GetByID(id string) (modelcinemacity.Response, error) {
	//TODO implement me
	panic("implement me")
}

func New(repo repositorycinemacity.RepositoryCinemaCity) ServiceCinemaCity {
	return &service{repo: repo}
}
