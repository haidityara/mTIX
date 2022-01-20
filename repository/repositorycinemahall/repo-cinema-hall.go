package repositorycinemahall

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
)

type RepositoryCinemaHall interface {
	Create(data entity.CinemaHall) (entity.CinemaHall, error)
	GetByID(id string) (entity.CinemaHall, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data entity.CinemaHall) (entity.CinemaHall, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.CinemaHall{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.CinemaHall, error) {
	var data entity.CinemaHall
	err := r.db.Preload("Cinema").
		Preload("CinemaSeat").
		Preload("Cinema.City").
		Where("id = ?", id).First(&data).Error
	if err != nil {
		return entity.CinemaHall{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryCinemaHall {
	return &repository{db}
}
