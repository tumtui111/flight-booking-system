package domain

import (
	"time"
)

type BookingHistory struct {
	FlightID     string    `json:"flight_id"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	Departure    time.Time `json:"departure"`
	BookingID    string    `json:"booking_id"`
	Seat         string    `json:"seat"`
	Price        float64   `json:"price"`
	Status       string    `json:"status"`
	RefundAmount float64   `json:"refund_amount,omitempty"`
}

type Passenger struct {
	PassengerID     string           `json:"passenger_id"`
	IsFrequentFlyer bool             `json:"is_frequent_flyer"`
	BookingHistory  []BookingHistory `json:"booking_history"`
}
