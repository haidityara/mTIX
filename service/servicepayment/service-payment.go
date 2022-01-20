package servicepayment

import (
	"github.com/haidityara/mtix/constant"
	"github.com/haidityara/mtix/entity"
	"github.com/haidityara/mtix/model/modelpayment"
	"github.com/haidityara/mtix/repository/repositorypayment"
	"github.com/jinzhu/copier"
)

type ServicePayment interface {
	Create(request modelpayment.Request) (modelpayment.ResponseStore, error)
	GetByID(id string) (modelpayment.ResponseGet, error)
}

type service struct {
	repo repositorypayment.RepositoryPayment
}

func (s *service) Create(request modelpayment.Request) (modelpayment.ResponseStore, error) {
	payment := entity.Payment{}
	copier.Copy(&payment, &request)
	response, err := s.repo.Create(payment)
	if err != nil {
		return modelpayment.ResponseStore{}, err
	}
	resp := modelpayment.ResponseStore{}
	copier.Copy(&resp, &response)
	paymentMethod := constant.PaymentMethods[response.PaymentMethod]
	resp.PaymentMethod = paymentMethod
	return resp, nil
}

func (s service) GetByID(id string) (modelpayment.ResponseGet, error) {
	response, err := s.repo.GetByID(id)
	if err != nil {
		return modelpayment.ResponseGet{}, err
	}
	resp := modelpayment.ResponseGet{}
	copier.Copy(&resp, &response)
	paymentMethod := constant.PaymentMethods[response.PaymentMethod]
	resp.PaymentMethod = paymentMethod
	return resp, nil
}

func New(repo repositorypayment.RepositoryPayment) ServicePayment {
	return &service{repo: repo}
}
