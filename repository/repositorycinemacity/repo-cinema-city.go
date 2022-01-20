package repositorycinemacity

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
)

type RepositoryCinemaCity interface {
	Create(data entity.CinemaCity) (entity.CinemaCity, error)
	GetByID(id string) (entity.CinemaCity, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data entity.CinemaCity) (entity.CinemaCity, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.CinemaCity{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.CinemaCity, error) {
	data := entity.CinemaCity{}
	err := r.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return entity.CinemaCity{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryCinemaCity {
	return &repository{db}
}
