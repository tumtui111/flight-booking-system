package constant

import (
	"errors"
)

var (
	ErrFlightNotFound    = errors.New("flight not found")
	ErrBookingNotFound   = errors.New("booking not found")
	ErrPassengerNotFound = errors.New("passenger not found")
	ErrNoAvailableSeat   = errors.New("no available seats found")
)
