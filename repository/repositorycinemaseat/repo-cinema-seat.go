package repositorycinemaseat

import (
	"github.com/haidityara/mtix/entity"
	"gorm.io/gorm"
)

type RepositoryCinemaSeat interface {
	Create(data entity.CinemaSeat) (entity.CinemaSeat, error)
	GetByID(id string) (entity.CinemaSeat, error)
	GetAvailableSeats(showID uint, cinemaHallID uint) ([]entity.CinemaSeat, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetByHall(id uint) ([]entity.CinemaSeat, error) {
	var cinemaSeats []entity.CinemaSeat
	err := r.db.Where("cinema_hall_id = ?", id).Find(&cinemaSeats).Error
	return cinemaSeats, err
}

func (r *repository) BookedSeat(showID uint, hallID uint) ([]entity.CinemaSeat, error) {
	var bookingDetail []entity.BookingDetail
	err := r.db.Preload("CinemaSeat", "cinema_hall_id = ?", hallID).
		Where("show_id = ?", showID).
		Find(&bookingDetail).Error

	var cinemaSeats []entity.CinemaSeat
	for _, book := range bookingDetail {
		cinemaSeats = append(cinemaSeats, *book.CinemaSeat)
	}

	return cinemaSeats, err
}

func (r *repository) GetAvailableSeats(showID uint, cinemaHallID uint) ([]entity.CinemaSeat, error) {

	bookedSeat, err := r.BookedSeat(showID, cinemaHallID)
	if err != nil {
		return []entity.CinemaSeat{}, err
	}
	seatHall, err := r.GetByHall(cinemaHallID)
	if err != nil {
		return []entity.CinemaSeat{}, err
	}

	var availableSeats []entity.CinemaSeat
	for _, seat := range seatHall {
		var isBooked bool
		for _, booked := range bookedSeat {
			if booked.ID == seat.ID {
				isBooked = true
				break
			}
		}
		if !isBooked {
			availableSeats = append(availableSeats, seat)
		}
	}

	return availableSeats, nil
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
