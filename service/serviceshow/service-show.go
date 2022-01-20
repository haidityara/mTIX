package serviceshow

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/helper"
	"github.com/haidityara/mtix/model/modelshow"
	"github.com/haidityara/mtix/repository/repositoryshow"
	"github.com/jinzhu/copier"
	"time"
)

type ServiceShow interface {
	Create(request modelshow.Request) (modelshow.ResponseStore, error)
	GetByID(id string) (modelshow.ResponseGet, error)
	GetActiveShow() ([]modelshow.ResponseGet, error)
	GetAll() ([]modelshow.ResponseGet, error)
}
type service struct {
	repo repositoryshow.RepositoryShow
}

func (s *service) GetAll() ([]modelshow.ResponseGet, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return []modelshow.ResponseGet{}, err
	}
	var response []modelshow.ResponseGet
	for _, d := range data {
		var res modelshow.ResponseGet
		copier.Copy(&res, &d)
		res.StartTime = helper.DateTOString(d.StartTime)
		res.EndTime = helper.DateTOString(d.EndTime)
		response = append(response, res)
	}
	return response, nil
}

func (s *service) Create(request modelshow.Request) (modelshow.ResponseStore, error) {
	show := entity.Show{}
	copier.Copy(&show, &request)
	show.DateShow = helper.ToDateFormatter(request.DateShow)
	show.StartTime = helper.ToTimeFormatter(request.StartTime)
	show.EndTime = helper.ToTimeFormatter(request.EndTime)

	response, err := s.repo.Create(show)
	if err != nil {
		return modelshow.ResponseStore{}, err
	}
	resp := modelshow.ResponseStore{}
	copier.Copy(&resp, &response)
	return resp, nil
}

func (s service) GetByID(id string) (modelshow.ResponseGet, error) {
	response, err := s.repo.GetByID(id)
	if err != nil {
		return modelshow.ResponseGet{}, err
	}
	resp := modelshow.ResponseGet{}
	copier.Copy(&resp, &response)
	resp.StartTime = helper.DateTOString(response.StartTime)
	resp.EndTime = helper.DateTOString(response.EndTime)
	return resp, nil
}

func (s service) GetActiveShow() ([]modelshow.ResponseGet, error) {
	date := time.Now()
	response, err := s.repo.GetActiveShow(date)
	if err != nil {
		return []modelshow.ResponseGet{}, err
	}
	var resp []modelshow.ResponseGet
	copier.Copy(&resp, &response)
	return resp, nil
}

func New(repo repositoryshow.RepositoryShow) ServiceShow {
	return &service{repo: repo}
}
