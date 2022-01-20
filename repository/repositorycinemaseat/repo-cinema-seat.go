package repositorycinemaseat

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
)

type RepositoryCinemaSeat interface {
	Create(data entity.CinemaSeat) (entity.CinemaSeat, error)
	GetByID(id string) (entity.CinemaSeat, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data entity.CinemaSeat) (entity.CinemaSeat, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.CinemaSeat{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.CinemaSeat, error) {
	var data entity.CinemaSeat
	err := r.db.Preload("CinemaHall").
		Preload("CinemaHall.Cinema").
		Preload("CinemaHall.Cinema.City").
		Where("id = ?", id).First(&data).Error
	if err != nil {
		return entity.CinemaSeat{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryCinemaSeat {
	return &repository{db}
}
