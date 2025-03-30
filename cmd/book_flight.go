package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"flight-book-system/constant"
	"flight-book-system/domain"

	bookingSvc "flight-book-system/service/booking"
)

func BookFlight(bookingService *bookingSvc.BookingService) {

	defer fmt.Println("====================================================================================")

	var passengerID, seatClassInput, flightID, isUpgradeClass string

	// input passenger ID
	fmt.Println("Enter Passenger ID:")
	_, err := fmt.Scan(&passengerID)
	if err != nil {
		fmt.Println("error invalid passengerID format.")
		return
	}
	passengerID = strings.TrimSpace(passengerID)

	// input flight ID
	fmt.Println("Enter Flight ID:")
	_, err = fmt.Scan(&flightID)
	if err != nil {
		fmt.Println("error invalid FlightID format.")
		return
	}
	flightID = strings.TrimSpace(flightID)

	// input seat class
	fmt.Println("Enter Seat Class (Economy, Business, First):")
	_, err = fmt.Scan(&seatClassInput)
	if err != nil {
		fmt.Println("error invalid seat class format.")
		return
	}
	seatClassInput = strings.TrimSpace(seatClassInput)
	seatClass := domain.SeatClass(seatClassInput)
	bookingDate := time.Now()

	// book seat
	booking, err := bookingService.BookSeat(passengerID, flightID, seatClass, bookingDate)
	if err != nil {

		// other error handle
		if !errors.Is(err, constant.ErrNoAvailableSeat) {
			fmt.Println("Booking failed:", err)
			return
		}

		// not have seat available error and upgrade to higher class logic
		classIndex := -1
		for i, class := range domain.ClassOrder {
			if class == seatClass {
				classIndex = i
				break
			}
		}

		// check invalid class or First class that cannot be upgraded
		if classIndex == -1 || classIndex == 3 {
			fmt.Println("Booking failed: ", err)
			return
		}

		fmt.Printf("Booking failed: No seats available in %v. Upgrade to %s? (yes/no)\n", seatClass, domain.ClassOrder[classIndex+1])
		_, err := fmt.Scan(&isUpgradeClass)
		if err != nil {
			fmt.Println("invalid seat class format.")
			return
		}

		// re-booking with higher class
		if isUpgradeClass == "yes" {
			seatClass = domain.ClassOrder[classIndex+1]
			booking, err := bookingService.BookSeat(passengerID, flightID, seatClass, bookingDate)
			if err != nil {
				fmt.Println("Booking failed: ", err)
			} else {
				b, _ := json.Marshal(booking)
				fmt.Printf("Booking successful: %+v\n", string(b))
			}

		}

	} else {

		b, _ := json.Marshal(booking)
		fmt.Printf("Booking successful: %+v\n", string(b))
	}
}
