package domain

import (
	"sync"
	"time"
)

type SeatClass string

const (
	NotAvailable SeatClass = "NotAvailable"
	Economy  SeatClass = "Economy"
	Business SeatClass = "Business"
	First    SeatClass = "First"
)

var ClassOrder = []SeatClass{NotAvailable, Economy, Business, First}

type SeatInfo struct {
	Total      int
	Available  int
	BasePrice  float64
	SeatMap   map[string]bool // for check seat available after booking cancelled
}

type Flight struct {
	FlightID   string
	Origin     string
	Destination string
	Departure  time.Time
	Seats      map[SeatClass]*SeatInfo
	ReservedSeats map[SeatClass]map[string]bool // for check seat available after booking cancelled
	Mutex      sync.Mutex
}

