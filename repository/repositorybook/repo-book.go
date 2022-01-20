package repositorybook

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
	"log"
)

type RepositoryBook interface {
	Create(data entity.Booking) (entity.Booking, error)
	GetByID(id string) (entity.Booking, error)
	GetByUserID(id string) (entity.Booking, error)
	GetAll() ([]entity.Booking, error)
	CreateDetailBook(data []entity.BookingDetail) ([]entity.BookingDetail, error)
	UpdateStatusBook(data entity.Booking) (entity.Booking, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) UpdateStatusBook(data entity.Booking) (entity.Booking, error) {
	err := r.db.Model(&data).Where("id = ?", data.ID).Update("status", data.Status).Error
	if err != nil {
		return entity.Booking{}, err
	}
	return data, nil
}

func (r *repository) CreateDetailBook(data []entity.BookingDetail) ([]entity.BookingDetail, error) {
	var temp []entity.BookingDetail
	for _, v := range data {
		log.Println(v)
		show := entity.Show{}
		if err := r.db.Where("id = ?", v.ShowID).First(&show).Error; err != nil {
			return []entity.BookingDetail{}, err
		}
		seat := entity.CinemaSeat{}
		if err := r.db.Where("id = ?", v.CinemaSeatID).First(&seat).Error; err != nil {
			return []entity.BookingDetail{}, err
		}
		v.NumberOFSeat = seat.SeatNum
		v.Price = show.Price
		temp = append(temp, v)
	}
	return temp, nil
}

func (r *repository) Create(data entity.Booking) (entity.Booking, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Booking{}, err
	}
	return data, nil
}

func (r repository) GetByID(id string) (entity.Booking, error) {
	var data entity.Booking
	err := r.db.Where("id = ?", id).
		Preload("BookingDetail").Preload("User").
		First(&data).Error
	if err != nil {
		return entity.Booking{}, err
	}
	return data, nil
}

func (r repository) GetByUserID(id string) (entity.Booking, error) {
	var data entity.Booking
	err := r.db.Where("user_id = ?", id).First(&data).Error
	if err != nil {
		return entity.Booking{}, err
	}
	return data, nil
}

func (r repository) GetAll() ([]entity.Booking, error) {
	var data []entity.Booking
	err := r.db.Find(&data).Error
	if err != nil {
		return []entity.Booking{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryBook {
	return &repository{db}
}
