package repositorymovie

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
)

type RepositoryMovie interface {
	Create(data entity.Movie) (entity.Movie, error)
	GetByID(id string) (entity.Movie, error)
	GetAll() ([]entity.Movie, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data entity.Movie) (entity.Movie, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Movie{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.Movie, error) {
	var data entity.Movie
	err := r.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		return entity.Movie{}, err
	}
	return data, nil
}

func (r *repository) GetAll() ([]entity.Movie, error) {
	var data []entity.Movie
	err := r.db.Find(&data).Error
	if err != nil {
		return []entity.Movie{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryMovie {
	return &repository{db}
}
