package repository

import (
	"flight-book-system/domain"
	"fmt"
	"sync"
)

type PassengerRepository struct {
	Passengers map[string]*domain.Passenger
	Mutex      sync.Mutex
}

func NewPassengerRepository() *PassengerRepository {
	return &PassengerRepository{
		Passengers: make(map[string]*domain.Passenger),
	}
}

func (pr *PassengerRepository) AddPassenger(passenger *domain.Passenger) {
	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	pr.Passengers[passenger.PassengerID] = passenger
}

func (pr *PassengerRepository) GetPassenger(passengerID string) (*domain.Passenger, bool) {
	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	passenger, exists := pr.Passengers[passengerID]
	return passenger, exists
}

func (pr *PassengerRepository) UpdatePassengerBookingStatus(booking *domain.Booking, status string) {

	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	passenger, exists := pr.Passengers[booking.PassengerID]
	if !exists {
		fmt.Println("not found passenger id to update status")
		return
	}

	for i, history := range passenger.BookingHistory {
		if history.FlightID == booking.FlightID && history.Seat == booking.Seat {
			passenger.BookingHistory[i].Status = status
		}
	}
}

func (pr *PassengerRepository) UpdatePassengerBookingRefundAmount(booking *domain.Booking, refundAmount float64) {

	pr.Mutex.Lock()
	defer pr.Mutex.Unlock()

	passenger, exists := pr.Passengers[booking.PassengerID]
	if !exists {
		fmt.Println("not found passenger id to update status")
		return
	}

	for i, history := range passenger.BookingHistory {
		if history.FlightID == booking.FlightID && history.Seat == booking.Seat {
			passenger.BookingHistory[i].RefundAmount = refundAmount
		}
	}
}
