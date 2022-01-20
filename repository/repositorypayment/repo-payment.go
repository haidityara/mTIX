package repositorypayment

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
	"strconv"
)

type RepositoryPayment interface {
	Create(data entity.Payment) (entity.Payment, error)
	GetByID(id string) (entity.Payment, error)
	GetByUserID(id string) (entity.Payment, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) UpdateStatusBook(bookID uint, status int) error {
	return r.db.Where("id = ?", bookID).First(&entity.Booking{}).Update("status", status).Error
}

func (r *repository) CountAmount(bookID uint) (float64, error) {
	var booksDetails []entity.BookingDetail
	err := r.db.Where("booking_id = ?", bookID).Find(&booksDetails).Error
	if err != nil {
		return 0, err
	}
	amount := 0.0
	for _, book := range booksDetails {
		price, _ := strconv.ParseFloat(book.Price, 64)
		amount += price
	}
	return amount, nil
}

func (r *repository) Create(data entity.Payment) (entity.Payment, error) {

	amount, err := r.CountAmount(data.BookingID)
	if err != nil {
		return entity.Payment{}, err
	}
	data.Amount = strconv.FormatFloat(amount, 'f', -1, 64)

	err = r.db.Create(&data).Error
	if err != nil {
		return entity.Payment{}, err
	}
	// status 1 = payment request by user
	err = r.UpdateStatusBook(data.BookingID, 1)
	if err != nil {
		return entity.Payment{}, err
	}
	return data, nil
}

func (r *repository) GetByID(id string) (entity.Payment, error) {
	var data entity.Payment
	err := r.db.Where("id = ?", id).
		Preload("Booking").Preload("Booking.User").Preload("Booking.BookingDetail").
		First(&data).Error
	if err != nil {
		return entity.Payment{}, err
	}

	return data, nil
}

func (r repository) GetByUserID(id string) (entity.Payment, error) {
	var data entity.Payment
	err := r.db.Where("user_id = ?", id).First(&data).Error
	if err != nil {
		return entity.Payment{}, err
	}
	return data, nil
}

func New(db *gorm.DB) RepositoryPayment {
	return &repository{db}
}
