package constant

import (
	"fmt"
)

var (
	ERR_FLIGHT_NOT_FOUND    = fmt.Errorf("flight not found")
	ERR_BOOKING_NOT_FOUND   = fmt.Errorf("booking not found")
	ERR_PASSENGER_NOT_FOUND = fmt.Errorf("passenger not found")
	ERR_NO_AVAILABLE_SEAT   = fmt.Errorf("no available seats found")
)
