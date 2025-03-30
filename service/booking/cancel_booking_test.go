package service

/*
Cancel
1. cancel with full refund
2. cancel with refund including cancellation fee
3. cancel booking id not found
4. re-cancel for cancelled booking
3. TestCancelThenReleaseBookingSeatToAnotherPassenger (Assert Value) [scenario]

*/

import (
	mockFlight "flight-book-system/mocks/service/flight"
	mockPassenger "flight-book-system/mocks/service/passenger"
	"fmt"

	"flight-book-system/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func SetupTestCancelBooking() (*mockFlight.IFlightRepository, *mockPassenger.IPassengerRepository, *BookingService, domain.Flight) {
	var (
		mockFlight    = new(mockFlight.IFlightRepository)
		mockPassenger = new(mockPassenger.IPassengerRepository)
		service       = NewBookingService(mockFlight, mockPassenger)

		flightInfo = domain.Flight{
			FlightID:    "AB123",
			Origin:      "JFK",
			Destination: "LAX",
			Departure:   time.Now().AddDate(0, 0, 10),
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
	)

	return mockFlight, mockPassenger, service, flightInfo
}

func Test_Cancel_Booking(t *testing.T) {
	t.Run("Success case | cancel booking seat | got full refund (class = Economy)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo := SetupTestCancelBooking()

		booking := &domain.Booking{
			BookingID:   "B1",
			PassengerID: "TEST01",
			FlightID:    "AB123",
			Seat:        "1A",
			Price:       300.0,
			Class:       "Economy",
			Status:      "Confirmed",
		}

		BookService.Bookings["B1"] = booking
		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)

		fmt.Println(flightInfo.Departure)

		mockPassenger.On("UpdatePassengerBookingStatus", booking, "Cancelled").Return()
		mockPassenger.On("UpdatePassengerBookingRefundAmount", booking, 300.0).Return()

		resp, err := BookService.CancelBooking("B1")

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 300.0, resp.RefundAmount)
		assert.Equal(t, "Cancelled", resp.Status)
	})

	t.Run("Success case | cancel booking seat | got refund deduct with cancellation fee (class = Business)", func(t *testing.T) {

		mockFlight, mockPassenger, BookService, flightInfo := SetupTestCancelBooking()

		flightInfo.Departure = time.Now().AddDate(0,0,1)

		booking := &domain.Booking{
			BookingID:   "B1",
			PassengerID: "TEST01",
			FlightID:    "AB123",
			Seat:        "1A",
			Price:       1000.0,
			Class:       "Business",
			Status:      "Confirmed",
		}

		BookService.Bookings["B1"] = booking
		mockFlight.On("GetFlight", "AB123").Return(&flightInfo, true)

		fmt.Println(flightInfo.Departure)

		mockPassenger.On("UpdatePassengerBookingStatus", booking, "Cancelled").Return()
		mockPassenger.On("UpdatePassengerBookingRefundAmount", booking, 900.0).Return()

		resp, err := BookService.CancelBooking("B1")

		assert.Nil(t, err)
		assert.Equal(t, "B1", resp.BookingID)
		assert.Equal(t, 900.0, resp.RefundAmount)
		assert.Equal(t, "Cancelled", resp.Status)
	})

	t.Run("Fail case | cancel booking seat | booking id not found", func(t *testing.T) {

		_, _, BookService, _ := SetupTestCancelBooking()
		_, err := BookService.CancelBooking("B1")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "booking not found")
	})

	t.Run("Fail case | cancel booking seat | re-cancelled for booking status already cancelled", func(t *testing.T) {

		_, _, BookService, _ := SetupTestCancelBooking()

		booking := &domain.Booking{
			BookingID:   "B1",
			PassengerID: "TEST01",
			FlightID:    "AB123",
			Seat:        "1A",
			Price:       1000.0,
			Class:       "Business",
			Status:      "Cancelled",
			RefundAmount: 1000.0,
		}

		BookService.Bookings["B1"] = booking

		_, err := BookService.CancelBooking("B1")

		assert.NotNil(t, err)
		assert.EqualError(t, err, "BookingID: B1 already cancelled")
	})

}
