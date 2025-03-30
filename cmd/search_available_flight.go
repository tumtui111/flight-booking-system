package cmd

import (
	flightRepo "flight-book-system/repository/flight"
	"fmt"
	"time"
)

func SearchAvailableFlight(flightRepo *flightRepo.FlightRepository) {

	defer fmt.Println("====================================================================================")

	var origin, destination, dateInput string

	fmt.Println("Enter Origin:")
	_, err := fmt.Scan(&origin)
	if err != nil {
		fmt.Println("invalid flight ID format.")
		return
	}

	fmt.Println("Enter Destination:")
	_, err = fmt.Scan(&destination)
	if err != nil {
		fmt.Println("invalid flight ID format.")
		return
	}

	fmt.Println("Enter Date (YYYY-MM-DD):")
	_, err = fmt.Scan(&dateInput)
	if err != nil {
		fmt.Println("invalid flight ID format.")
		return
	}

	date, err := time.Parse("2006-01-02", dateInput)
	if err != nil {
		fmt.Println("Invalid date format.")
		return
	}

	flights := flightRepo.SearchFlights(origin, destination, date)
	if len(flights) == 0 {
		fmt.Println("No available flights found.")
	} else {
		for _, f := range flights {
			fmt.Printf("Flight ID: %s, Departure: %s\n", f.FlightID, f.Departure)
		}
	}
}
