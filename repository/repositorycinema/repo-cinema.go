package repositorycinema

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
)

type RepositoryCinema interface {
	Create(data entity.Cinema) (entity.Cinema, error)
	GetByID(id string) (entity.Cinema, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data entity.Cinema) (entity.Cinema, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Cinema{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.Cinema, error) {
	cinema := entity.Cinema{}
	err := r.db.Preload("City").Where("id = ?", id).First(&cinema).Error
	if err != nil {
		return entity.Cinema{}, err
	}
	return cinema, nil
}

func New(db *gorm.DB) RepositoryCinema {
	return &repository{db}
}
