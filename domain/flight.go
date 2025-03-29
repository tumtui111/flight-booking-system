package domain

import (
	"sync"
	"time"
)

type SeatClass string

const (
	NotAvailable SeatClass = "NotAvailable"
	Economy      SeatClass = "Economy"
	Business     SeatClass = "Business"
	First        SeatClass = "First"
)

var ClassOrder = []SeatClass{NotAvailable, Economy, Business, First}

type SeatInfo struct {
	Total     int             `json:"total"`
	Available int             `json:"available"`
	BasePrice float64         `json:"base_price"`
	SeatMap   map[string]bool `json:"seat_map,omitempty"` // for check seat available after booking cancelled
}

type Flight struct {
	FlightID      string                        `json:"flight_id"`
	Origin        string                        `json:"origin"`
	Destination   string                        `json:"destination"`
	Departure     time.Time                     `json:"departure"`
	Seats         map[SeatClass]*SeatInfo       `json:"seats"`
	ReservedSeats map[SeatClass]map[string]bool `json:"reserved_seats"` // for check seat available after booking cancelled
	Mutex         sync.Mutex                    `json:"mutex,omitempty"`
}
