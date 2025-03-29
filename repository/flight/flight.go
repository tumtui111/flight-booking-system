package repository

import (
	"flight-book-system/domain"
	"sync"
	"time"
	"fmt"
)

type FlightRepository struct {
	Flights map[string]*domain.Flight
	Mutex   sync.Mutex
}

func NewFlightRepository() *FlightRepository {
	return &FlightRepository{
		Flights: make(map[string]*domain.Flight),
	}
}

func (fr *FlightRepository) AddFlight(flight *domain.Flight) {
	fr.Mutex.Lock()
	defer fr.Mutex.Unlock()

	//check if flightID already exists
	if  fr.Flights[flight.FlightID] != nil {
		fmt.Println("FlightID already exists.")
		return 
	}

	fr.Flights[flight.FlightID] = flight
}

func (fr *FlightRepository) GetFlight(flightID string) (*domain.Flight, bool) {
	fr.Mutex.Lock()
	defer fr.Mutex.Unlock()
	flight, exists := fr.Flights[flightID]
	return flight, exists
}

func (fr *FlightRepository) SearchFlights(origin, destination string, date time.Time) []*domain.Flight {
	fr.Mutex.Lock()
	defer fr.Mutex.Unlock()
	
	var availableFlights []*domain.Flight
	for _, flight := range fr.Flights {
		if flight.Origin == origin && flight.Destination == destination && flight.Departure.Format("2006-01-02") == date.Format("2006-01-02") {
			availableFlights = append(availableFlights, flight)
		}
	}
	return availableFlights
}

