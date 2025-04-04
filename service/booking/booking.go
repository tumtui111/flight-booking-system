package service

import (
	"flight-book-system/constant"
	"flight-book-system/domain"

	"fmt"
	"math"
	"sync"
	"time"
)

type IFlightRepository interface {
	GetFlight(flightID string) (*domain.Flight, bool)
}

type IPassengerRepository interface {
	AddPassenger(passenger *domain.Passenger)
	GetPassenger(passengerID string) (*domain.Passenger, bool)
	UpdatePassengerBookingStatus(booking *domain.Booking, status string)
	UpdatePassengerBookingRefundAmount(booking *domain.Booking, refundAmount float64)
}

type BookingService struct {
	FlightRepo    IFlightRepository
	PassengerRepo IPassengerRepository
	Bookings      map[string]*domain.Booking
	Passengers    map[string]*domain.Passenger
	Mutex         sync.Mutex
}

func NewBookingService(flightRepo IFlightRepository, passengerRepo IPassengerRepository) *BookingService {
	return &BookingService{
		FlightRepo:    flightRepo,
		PassengerRepo: passengerRepo,
		Bookings:      make(map[string]*domain.Booking),
		Passengers:    make(map[string]*domain.Passenger),
	}
}

func (bs *BookingService) BookSeat(passengerID, flightID string, seatClass domain.SeatClass, bookingDate time.Time) (*domain.Booking, error) {
	bs.Mutex.Lock()
	defer bs.Mutex.Unlock()

	// flight detail update
	flight, exists := bs.FlightRepo.GetFlight(flightID)
	if !exists {
		return nil, constant.ErrFlightNotFound
	}

	flight.Mutex.Lock()
	defer flight.Mutex.Unlock()

	// check seat available -> then return to ask for upgrade class
	seatInfo, exists := flight.Seats[seatClass]
	if !exists || seatInfo.Available == 0 {
		return nil, constant.ErrNoAvailableSeat
	}

	// check passenger detail exists -> if not, created
	passenger, exists := bs.PassengerRepo.GetPassenger(passengerID)
	if !exists {
		// create new passenger detail
		passenger = &domain.Passenger{
			PassengerID:     passengerID,
			IsFrequentFlyer: false,
			BookingHistory:  []domain.BookingHistory{}}

		bs.PassengerRepo.AddPassenger(passenger)
	}

	// dynamic pricing ( think before available seat is decreased)
	price := calculateDynamicPricing(seatInfo.BasePrice, flight, bookingDate, seatInfo, passenger)

	//logic for handling seatID in each seat class
	var assignedSeat string
	for row := 'A'; row <= 'Z'; row++ {
		for col := 1; col <= 6; col++ {
			seat := fmt.Sprintf("%d%c", col, row)
			if !flight.ReservedSeats[seatClass][seat] {
				assignedSeat = seat
				flight.ReservedSeats[seatClass][seat] = true
				seatInfo.Available--
				break
			}
		}
		if assignedSeat != "" {
			break
		}
	}

	if assignedSeat == "" {
		return nil, constant.ErrNoAvailableSeat
	}

	//booking detail update
	booking := &domain.Booking{
		BookingID:   fmt.Sprintf("B%d", len(bs.Bookings)+1),
		PassengerID: passengerID,
		FlightID:    flightID,
		Seat:        assignedSeat,
		Price:       price,
		Class:       seatClass,
		Status:      constant.STATUS_CONFIRMED,
	}
	bs.Bookings[booking.BookingID] = booking

	// keep passenger booking history
	passenger.BookingHistory = append(passenger.BookingHistory, domain.BookingHistory{
		FlightID:    flightID,
		Origin:      flight.Origin,
		Destination: flight.Destination,
		Departure:   flight.Departure,
		BookingID:   booking.BookingID,
		BookingDate: bookingDate,
		Class:       seatClass,
		Seat:        booking.Seat,
		Price:       booking.Price,
		Status:      booking.Status,
	})

	return booking, nil
}

func (bs *BookingService) CancelBooking(bookingID string) (*domain.Booking, error) {
	bs.Mutex.Lock()
	defer bs.Mutex.Unlock()

	booking, exists := bs.Bookings[bookingID]
	if !exists {
		return nil, constant.ErrBookingNotFound
	}

	// check if already cancelled, cannot cancelled again
	if booking.Status == constant.STATUS_CANCELLED {
		return nil, fmt.Errorf("BookingID: %s already cancelled", bookingID)
	}

	flight, flightExists := bs.FlightRepo.GetFlight(booking.FlightID)
	if !flightExists {
		return nil, constant.ErrFlightNotFound
	}

	flight.Mutex.Lock()
	delete(flight.ReservedSeats[booking.Class], booking.Seat)
	flight.Seats[booking.Class].Available++
	flight.Seats[booking.Class].SeatMap[booking.Seat] = false
	defer flight.Mutex.Unlock()

	// Calculate refund amount
	today := time.Now()
	daysBeforeDeparture := int(flight.Departure.Sub(today).Hours() / 24)
	refund := booking.Price
	fee := 0.0

	// if last minute cancellation, have cancellation fee 10% (refund = refund - cancellation fee)
	if daysBeforeDeparture < 7 {
		fee = booking.Price * 0.10
		refund -= fee
	}

	booking.Status = constant.STATUS_CANCELLED
	booking.RefundAmount = refund

	//update status and refund amount in passenger booking history
	bs.PassengerRepo.UpdatePassengerBookingStatus(booking, constant.STATUS_CANCELLED)
	bs.PassengerRepo.UpdatePassengerBookingRefundAmount(booking, refund)

	return booking, nil
}

func calculateDynamicPricing(basePrice float64, flight *domain.Flight, bookingDate time.Time, seatInfo *domain.SeatInfo, passenger *domain.Passenger) float64 {

	// init base price
	price := basePrice

	// calculate up to date booking
	daysBeforeDeparture := int(flight.Departure.Sub(bookingDate).Hours() / 24)
	if daysBeforeDeparture > 30 {
		//early booking
		price *= 0.9 // 10% discount
	} else if daysBeforeDeparture < 7 {
		//last minute condition
		price *= 1.2 // 20% surcharge
	}

	// More seats booked → increase price (1% per 10% seats booked tier)
	bookedCount := seatInfo.Total - seatInfo.Available
	tier := int(math.Floor(float64(bookedCount) / float64(seatInfo.Total) * 10)) // not increase for first 10 %
	price *= 1.0 + float64(tier)*0.01

	// check is frequent flyer discount 10% for frequent flyer
	if passenger.IsFrequentFlyer {
		price *= 0.90
	}

	return price
}
