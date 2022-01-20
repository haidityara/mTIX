package servicemovie

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelmovie"
	"github.com/haidityara/mtix/repository/repositorymovie"
	"github.com/haidityara/mtix/validation"
	"github.com/jinzhu/copier"
	"time"
)

type ServiceMovie interface {
	Create(request modelmovie.Request) (modelmovie.Response, error)
	GetByID(id string) (modelmovie.Response, error)
	GetAll() ([]modelmovie.Response, error)
}

type service struct {
	repo repositorymovie.RepositoryMovie
}

func (s *service) Create(request modelmovie.Request) (modelmovie.Response, error) {
	err := validation.ValidateMovieStore(request)
	if err != nil {
		return modelmovie.Response{}, err
	}
	movie := entity.Movie{}
	copier.Copy(&movie, &request)
	format := "2006-01-02"
	date, _ := time.Parse(format, request.RealiseDate)
	movie.RealiseDate = date
	create, err := s.repo.Create(movie)
	if err != nil {
		return modelmovie.Response{}, err
	}
	resp := modelmovie.Response{}
	copier.Copy(&resp, &create)
	return resp, nil
}

func (s *service) GetByID(id string) (modelmovie.Response, error) {
	get, err := s.repo.GetByID(id)
	if err != nil {
		return modelmovie.Response{}, err
	}
	resp := modelmovie.Response{}
	copier.Copy(&resp, &get)
	return resp, nil
}

func (s *service) GetAll() ([]modelmovie.Response, error) {
	get, err := s.repo.GetAll()
	if err != nil {
		return []modelmovie.Response{}, err
	}
	var resp []modelmovie.Response
	copier.Copy(&resp, &get)
	return resp, nil
}

func New(repo repositorymovie.RepositoryMovie) ServiceMovie {
	return &service{repo: repo}
}
