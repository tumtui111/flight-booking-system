package service

/*
Test case
	Booking
	1.dynamic pricing
		1.1 standard booking
		1.2 early booking
		1.3 last-minute booking
		1.4 standard booking and have 10% booked seat
		1.5 early booking and 10% booked seat
		1.6 last-minute booking and have 10% booked seat
		1.7 last-minute booking and have 10% booked seat with frequent flyer discount
	2. flight not found
	3. not available seat
	4. scenario case flight not available but have passenger cancel booking then book again
*/

import (
	mockFlight "flight-book-system/mocks/service/flight"
	mockPassenger "flight-book-system/mocks/service/passenger"

	"flight-book-system/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func SetupTestBooking() (*mockFlight.IFlightRepository, *mockPassenger.IPassengerRepository, *BookingService, domain.Flight, domain.Passenger, domain.Passenger) {
	var (
		mockFlight    = new(mockFlight.IFlightRepository)
		mockPassenger = new(mockPassenger.IPassengerRepository)
		service       = NewBookingService(mockFlight, mockPassenger)

		flightInfo = domain.Flight{
			FlightID:    "AB123",
			Origin:      "JFK",
			Destination: "LAX",
			Departure:   time.Date(2024, 7, 10, 8, 0, 0, 0, time.Local), // day 10 | month 7 | year 2024
			Seats: map[domain.SeatClass]*domain.SeatInfo{
				domain.NotAvailable: {Total: 10, Available: 9, BasePrice: 100, SeatMap: make(map[string]bool)}, // for test
				domain.Economy:      {Total: 100, Available: 100, BasePrice: 300, SeatMap: make(map[string]bool)},
				domain.Business:     {Total: 30, Available: 30, BasePrice: 1000, SeatMap: make(map[string]bool)},
				domain.First:        {Total: 10, Available: 0, BasePrice: 3000, SeatMap: make(map[string]bool)}, // test not available seat
			},
			ReservedSeats: map[domain.SeatClass]map[string]bool{
				domain.NotAvailable: make(map[string]bool),
				domain.Economy:      make(map[string]bool),
				domain.Business:     make(map[string]bool),
				domain.First:        make(map[string]bool),
			},
		}
		passengerDetailNew = domain.Passenger{
			PassengerID:     "TEST01",
			IsFrequentFlyer: false,
			BookingHistory:  []domain.BookingHistory{},
		}

		passengerDetailExists = domain.Passenger{
			PassengerID:     "TEST01",
			IsFrequentFlyer: true,
			BookingHistory: []domain.BookingHistory{
				{
					FlightID:    "AB123",
					Origin:      "JFK",
					Destination: "LAX",
					Departure:   time.Date(2024, 7, 10, 8, 0, 0, 0, time.Local),
					BookingID:   "B1",
					BookingDate: time.Date(2024, 7, 1, 8, 0, 0, 0, time.Local),
					Class:       "Economy",
					Seat:        "1A",
					Price:       300,
					Status:      "Confirmed",
				},
			},
		}
	)

	return mockFlight, mockPassenger, service, flightInfo, passengerDetailNew, passengerDetailExists
}

func Test_Booking(t *testing.T) {

	t.Run("Success case | Book seat | new passenger and standard booking (class = Economy)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, passengerDetailNew, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(nil, false)
		mockPassenger.On("AddPassenger", &passengerDetailNew).Return(nil)

		resp, err := BookService.BookSeat("TEST01", "AB123", "Economy", time.Date(2024, 7, 1, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 300.0, resp.Price)
	})

	t.Run("Success case | Book seat | new passenger and early booking (class = Economy)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, passengerDetailNew, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(nil, false)
		mockPassenger.On("AddPassenger", &passengerDetailNew).Return(nil)

		resp, err := BookService.BookSeat("TEST01", "AB123", "Economy", time.Date(2024, 6, 1, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 270.0, resp.Price)
	})

	t.Run("Success case | Book seat | new passenger and last-minute booking (class = Economy)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, passengerDetailNew, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(nil, false)
		mockPassenger.On("AddPassenger", &passengerDetailNew).Return(nil)

		resp, err := BookService.BookSeat("TEST01", "AB123", "Economy", time.Date(2024, 7, 8, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 360.0, resp.Price)
	})

	t.Run("Success case | Book seat | new passenger and standard booking with 10% booked seat (class = NotAvailabe)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, passengerDetailNew, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(nil, false)
		mockPassenger.On("AddPassenger", &passengerDetailNew).Return(nil)

		resp, err := BookService.BookSeat("TEST01", "AB123", "NotAvailable", time.Date(2024, 7, 1, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 101.0, resp.Price)
	})

	t.Run("Success case | Book seat | new passenger and early booking with 10% booked seat (class = NotAvailabe)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, passengerDetailNew, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(nil, false)
		mockPassenger.On("AddPassenger", &passengerDetailNew).Return(nil)

		resp, err := BookService.BookSeat("TEST01", "AB123", "NotAvailable", time.Date(2024, 6, 1, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 90.9, resp.Price)
	})

	t.Run("Success case | Book seat | new passenger and last-minute booking with 10% booked seat (class = NotAvailabe)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, passengerDetailNew, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(nil, false)
		mockPassenger.On("AddPassenger", &passengerDetailNew).Return(nil)

		resp, err := BookService.BookSeat("TEST01", "AB123", "NotAvailable", time.Date(2024, 7, 8, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 121.2, resp.Price)
	})

	t.Run("Success case | Book seat | existing passenger with IsFrequentFlyer and last-minute booking with 10% booked seat (class = NotAvailabe)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo, _, passengerDetailExists := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)
		mockPassenger.On("GetPassenger", "TEST01").Return(&passengerDetailExists, true)

		resp, err := BookService.BookSeat("TEST01", "AB123", "NotAvailable", time.Date(2024, 7, 8, 8, 0, 0, 0, time.Local))

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 109.08, resp.Price)

		// base price 														= 100
		// last minute booking (add 20% from base price) 					= 20
		// 10% booked seat (add 1% from base price + last minute booking) 	= 1.2
		// 																	= 100 + 20 + 1.2 = 121.2
		// existing passenger frequent flyer = true, discount 10% 			= 121.2 - 12.12 = 109.08 //
	})

	t.Run("Fail case | Book seat | flight not found", func(t *testing.T) {

		mockFlight, _, BookService, _, _, _ := SetupTestBooking()
		mockFlight.On("GetFlight", "AB123").Return(nil, false)

		_, err := BookService.BookSeat("TEST01", "AB123", "First", time.Date(2024, 7, 8, 8, 0, 0, 0, time.Local))

		assert.NotNil(t, err)
		assert.EqualError(t, err, "flight not found")
	})

	t.Run("Fail case | Book seat | seat not availble (class = First)", func(t *testing.T) {

		mockFlight, _, BookService, flightInfo, _, _ := SetupTestBooking()

		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)

		_, err := BookService.BookSeat("TEST01", "AB123", "First", time.Date(2024, 7, 8, 8, 0, 0, 0, time.Local))

		assert.NotNil(t, err)
		assert.EqualError(t, err, "no seats available in First")

	})

	t.Run("Success case | scenario case | Not available seat then cancel seat and booking again ", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, _, _, _ := SetupTestBooking()

		flightInfo := domain.Flight{
			FlightID:    "CD123",
			Origin:      "JFK",
			Destination: "LAX",
			Departure:   time.Now().AddDate(0, 0, 10),
			Seats: map[domain.SeatClass]*domain.SeatInfo{
				domain.Economy: {Total: 100, Available: 0, BasePrice: 300, SeatMap: make(map[string]bool)},
			},
			ReservedSeats: map[domain.SeatClass]map[string]bool{
				domain.Economy: make(map[string]bool),
			},
		}

		booking := &domain.Booking{
			BookingID:   "B1",
			PassengerID: "TEST01",
			FlightID:    "CD123",
			Seat:        "1A",
			Price:       300.0,
			Class:       "Economy",
			Status:      "Confirmed",
		}

		// cancel booking
		BookService.Bookings["B1"] = booking
		mockFlight.On("GetFlight", "CD123").Return(&flightInfo, true)
		mockPassenger.On("UpdatePassengerBookingStatus", booking, "Cancelled").Return()
		mockPassenger.On("UpdatePassengerBookingRefundAmount", booking, 300.0).Return()

		cancelResp, err := BookService.CancelBooking("B1")
		assert.Nil(t, err)
		assert.Equal(t, "Cancelled", cancelResp.Status)
		assert.Equal(t, 300.0, cancelResp.RefundAmount)           // cancel before flight 10 day, not include cancellation fee
		assert.Equal(t, "1A", cancelResp.Seat) 
		assert.Equal(t, 1, flightInfo.Seats["Economy"].Available) // check flight info available seat = 1
		// ------------

		// booking seat again
		passengerDetailExists := domain.Passenger{
			PassengerID:     "TEST01",
			IsFrequentFlyer: false,
			BookingHistory: []domain.BookingHistory{
				{},
			},
		}

		mockFlight.On("GetFlight", "CD123").Return(&flightInfo, true).Once()
		mockPassenger.On("GetPassenger", "TEST01").Return(&passengerDetailExists, true)

		bookingResp, err := BookService.BookSeat("TEST01", "CD123", "Economy", time.Now()) // book before flight 10 day, price add only 90% booked seat
		assert.Equal(t, "B2", bookingResp.BookingID)
		assert.Equal(t, 327.0, bookingResp.Price)
		assert.Equal(t, "1A", bookingResp.Seat)
		assert.Equal(t, 0, flightInfo.Seats["Economy"].Available)
	})

}
