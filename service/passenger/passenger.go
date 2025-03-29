package service

import (
	"flight-book-system/domain"
	"sync"
	"fmt"
	PassengerRepository "flight-book-system/repository/passenger"
)

type PassengerService struct {
	PassengerRepo *PassengerRepository.PassengerRepository
	Mutex      sync.Mutex
}

func NewPassengerService(passengerRepo *PassengerRepository.PassengerRepository) *PassengerService {
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