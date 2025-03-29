package service

import (
	"flight-book-system/domain"
	"sync"
	"fmt"
)

type IPassengerRepository interface {
	GetPassenger(passengerID string) (*domain.Passenger, bool)
}

type PassengerService struct {
	PassengerRepo IPassengerRepository
	Mutex      sync.Mutex
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
		return nil, fmt.Errorf("passenger not found")
	}
	return passenger, nil
}