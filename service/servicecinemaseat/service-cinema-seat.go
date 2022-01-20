package servicecinemaseat

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelcinemaseat"
	"github.com/haidityara/mtix/repository/repositorycinemaseat"
	"github.com/jinzhu/copier"
)

type ServiceCinemaSeat interface {
	Create(request modelcinemaseat.Request) (modelcinemaseat.ResponseStore, error)
	GetByID(id string) (modelcinemaseat.ResponseGet, error)
}

type service struct {
	repo repositorycinemaseat.RepositoryCinemaSeat
}

func (s *service) Create(request modelcinemaseat.Request) (modelcinemaseat.ResponseStore, error) {
	cinemaSeat := entity.CinemaSeat{}
	copier.Copy(&cinemaSeat, &request)
	response, err := s.repo.Create(cinemaSeat)
	if err != nil {
		return modelcinemaseat.ResponseStore{}, err
	}
	resp := modelcinemaseat.ResponseStore{}
	copier.Copy(&resp, &response)
	return resp, nil
}

func (s *service) GetByID(id string) (modelcinemaseat.ResponseGet, error) {
	data, err := s.repo.GetByID(id)
	if err != nil {
		return modelcinemaseat.ResponseGet{}, err
	}
	resp := modelcinemaseat.ResponseGet{}
	copier.Copy(&resp, &data)
	return resp, nil
}

func New(repo repositorycinemaseat.RepositoryCinemaSeat) ServiceCinemaSeat {
	return &service{repo: repo}
}
