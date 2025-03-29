package service

import (
	"flight-book-system/domain"
	repository "flight-book-system/repository/flight"
	"sync"
	"fmt"
)

type FlightService struct {
	FlightRepo *repository.FlightRepository
	Bookings   map[string]*domain.Booking
	Mutex      sync.Mutex
}

func NewFlightService(flightRepo *repository.FlightRepository) *FlightService {
	return &FlightService{
		FlightRepo: flightRepo,
	}
}

func (fs *FlightService) GetFlightInfo(flightID string) {
	flight, exists := fs.FlightRepo.GetFlight(flightID)
	if !exists {
		fmt.Println("Flight not found.")
		return
	}
	fmt.Printf("Flight ID: %s\nOrigin: %s\nDestination: %s\nDeparture: %s\n", flight.FlightID, flight.Origin, flight.Destination, flight.Departure)
	fmt.Println("Seat Availability:")
	for class, info := range flight.Seats {
		fmt.Printf("%s - Total: %d, Available: %d, Base Price: %.2f\n", class, info.Total, info.Available, info.BasePrice)
	}
}