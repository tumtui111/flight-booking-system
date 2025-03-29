package service

import (
	"testing"
	"time"

	// "github.com/stretchr/testify/assert"
	"flight-book-system/domain"
	mockFlight "flight-book-system/mocks/service/flight"

	"github.com/stretchr/testify/assert"
)

func Test_Flight(t *testing.T) {
	t.Run("Success case | GetFlightInfo | found data", func(t *testing.T) {

		flightInfo := domain.Flight{
			FlightID:    "AB123",
			Origin:      "JFK",
			Destination: "LAX",
			Departure:   time.Now(),
			Seats: map[domain.SeatClass]*domain.SeatInfo{
				domain.Economy:  {Total: 100, Available: 100, BasePrice: 300, SeatMap: make(map[string]bool)},
				domain.Business: {Total: 30, Available: 30, BasePrice: 1000, SeatMap: make(map[string]bool)},
				domain.First:    {Total: 10, Available: 10, BasePrice: 3000, SeatMap: make(map[string]bool)},
			},
			ReservedSeats: map[domain.SeatClass]map[string]bool{
				domain.NotAvailable: make(map[string]bool),
				domain.Economy:      make(map[string]bool),
				domain.Business:     make(map[string]bool),
				domain.First:        make(map[string]bool),
			},
		}

		var (
			mockFlight    = mockFlight.NewIFlightRepository(t)
			FlightService = NewFlightService(mockFlight)
		)

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		resp, err := FlightService.GetFlightInfo("AB123")
		assert.Nil(t, err)
		assert.Equal(t, resp.FlightID, "AB123")
		assert.Equal(t, resp.Seats["Economy"].Available, 100)
		assert.Equal(t, resp.Seats["First"].BasePrice, 3000.0)
	})

	t.Run("Fail case | GetFlightInfo | not found data", func(t *testing.T) {

		var (
			mockFlight    = mockFlight.NewIFlightRepository(t)
			FlightService = NewFlightService(mockFlight)
		)

		mockFlight.On("GetFlight", "AB123").Return(nil, false)

		_, err := FlightService.GetFlightInfo("AB123")
		assert.NotNil(t, err)

	})
}
