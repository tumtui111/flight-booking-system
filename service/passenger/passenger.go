package service

import (
	"flight-book-system/constant"
	"flight-book-system/domain"
	"sync"
)

type IPassengerRepository interface {
	AddPassenger(passenger *domain.Passenger)
	GetPassenger(passengerID string) (*domain.Passenger, bool)
	UpdatePassengerBookingStatus(booking *domain.Booking, status string)
	UpdatePassengerBookingRefundAmount(booking *domain.Booking, refundAmount float64)
}

type PassengerService struct {
	PassengerRepo IPassengerRepository
	Mutex         sync.Mutex
}

func NewPassengerService(passengerRepo IPassengerRepository) *PassengerService {
	return &PassengerService{
		PassengerRepo: passengerRepo,
	}
}

func (ps *PassengerService) GetPassengerDetails(passengerID string) (*domain.Passenger, error) {
	ps.Mutex.Lock()
	defer ps.Mutex.Unlock()

	passenger, exists := ps.PassengerRepo.GetPassenger(passengerID)
	if !exists {
		return nil, constant.ErrPassengerNotFound
	}
	return passenger, nil
}
