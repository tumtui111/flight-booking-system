package service

import (
	"flight-book-system/domain"
	"fmt"
	"sync"
)

type IFlightRepository interface {
	GetFlight(flightID string) (*domain.Flight, bool)
}

type FlightService struct {
	FlightRepo IFlightRepository
	Bookings   map[string]*domain.Booking
	Mutex      sync.Mutex
}

func NewFlightService(flightRepo IFlightRepository) *FlightService {
	return &FlightService{
		FlightRepo: flightRepo,
	}
}

func (fs *FlightService) GetFlightInfo(flightID string) (*domain.Flight, error) {
	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()

	flight, exists := fs.FlightRepo.GetFlight(flightID)
	if !exists {
		return nil, fmt.Errorf("flight not found")
	}

	return flight, nil
}