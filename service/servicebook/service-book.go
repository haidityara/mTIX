package servicebook

import (
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelbook"
	"github.com/haidityara/mtix/repository/repositorybook"
	"github.com/jinzhu/copier"
)

type ServiceBook interface {
	Create(request modelbook.Request) (modelbook.ResponseStore, error)
	GetByID(id string) (modelbook.ResponseGet, error)
	GetByUserID(id string) (modelbook.ResponseGet, error)
}

type service struct {
	repo repositorybook.RepositoryBook
}

func (s *service) Create(request modelbook.Request) (modelbook.ResponseStore, error) {
	book := entity.Booking{}
	copier.Copy(&book, &request)

	detail, err := s.repo.CreateDetailBook(book.BookingDetail)

	book.BookingDetail = detail

	create, err := s.repo.Create(book)
	if err != nil {
		return modelbook.ResponseStore{}, err
	}

	resp := modelbook.ResponseStore{}
	copier.Copy(&resp, &create)

	return resp, nil
}

func (s service) GetByID(id string) (modelbook.ResponseGet, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return modelbook.ResponseGet{}, err
	}

	resp := modelbook.ResponseGet{}
	copier.Copy(&resp, &book)

	return resp, nil
}

func (s service) GetByUserID(id string) (modelbook.ResponseGet, error) {
	//TODO implement me
	panic("implement me")
}

func New(repo repositorybook.RepositoryBook) ServiceBook {
	return &service{repo: repo}

}
