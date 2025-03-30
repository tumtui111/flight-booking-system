package cmd

import (
	"fmt"

	flightSvc "flight-book-system/service/flight"
)

func GetFlightInformation(flightService *flightSvc.FlightService) {

	defer fmt.Println("====================================================================================")

	var flightID string

	fmt.Println("Enter Flight ID:")
	_, err := fmt.Scan(&flightID)
	if err != nil {
		fmt.Println("invalid flight ID format.")
		return
	}

	flight, err := flightService.GetFlightInfo(flightID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Flight ID: %s\nOrigin: %s\nDestination: %s\nDeparture: %s\n", flight.FlightID, flight.Origin, flight.Destination, flight.Departure)
		fmt.Println("Seat Availability:")
		for class, info := range flight.Seats {
			fmt.Printf("%s - Total: %d, Available: %d, Base Price: %.2f\n", class, info.Total, info.Available, info.BasePrice)
		}
	}
}
