package repositoryshow

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
	"time"
)

type RepositoryShow interface {
	Create(data entity.Show) (entity.Show, error)
	GetByID(id string) (entity.Show, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data entity.Show) (entity.Show, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Show{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.Show, error) {
	var data entity.Show
	err := r.db.Where("id = ?", id).
		Preload("Movie").
		Preload("CinemaHall").Preload("CinemaHall.Cinema").Preload("CinemaHall.Cinema.City").
		First(&data).Error
	if err != nil {
		return entity.Show{}, err
	}
	return data, nil
}

func (r repository) GetActiveShow(date time.Time) ([]entity.Show, error) {
	var data []entity.Show
	err := r.db.Where("date > ?", date).Find(&data).Error
	if err != nil {
		return []entity.Show{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryShow {
	return &repository{db}
}
