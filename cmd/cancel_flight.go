package cmd

import (
	"encoding/json"
	"fmt"

	bookingSvc "flight-book-system/service/booking"
)

func CancelFlight(bookingService *bookingSvc.BookingService) {

	defer fmt.Println("====================================================================================")

	var bookingID string

	fmt.Println("Enter Booking ID:")
	_, err := fmt.Scan(&bookingID)
	if err != nil {
		fmt.Println("invalid booking ID format.")
		return
	}

	canceledBooking, err := bookingService.CancelBooking(bookingID)
	if err != nil {
		fmt.Println("Cancellation failed:", err)
	} else {
		b, _ := json.Marshal(canceledBooking)
		fmt.Printf("Booking cancelled: %+v\n", string(b))
	}
}
