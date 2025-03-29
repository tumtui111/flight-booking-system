/*
Copyright © 2025 worawit.l.official@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"flight-book-system/domain"
	flightRepo "flight-book-system/repository/flight"
	passengerRepo "flight-book-system/repository/passenger"
	bookingSvc "flight-book-system/service/booking"
	flightSvc "flight-book-system/service/flight"
	passengerSvc "flight-book-system/service/passenger"

	"github.com/spf13/cobra"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flight-book-system.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flight-book-system",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains`,

	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		var (
			flightRepo       = flightRepo.NewFlightRepository()
			passengerRepo    = passengerRepo.NewPassengerRepository()
			flightService    = flightSvc.NewFlightService(flightRepo)
			passengerService = passengerSvc.NewPassengerService(passengerRepo)
			bookingService   = bookingSvc.NewBookingService(flightRepo, passengerRepo)
		)

		InteractiveProgram(flightService, flightRepo, passengerService, passengerRepo, bookingService)
	},
}

func InteractiveProgram(
	flightService *flightSvc.FlightService, flightRepo *flightRepo.FlightRepository,
	passengerService *passengerSvc.PassengerService, passengerRepo *passengerRepo.PassengerRepository,
	bookingService *bookingSvc.BookingService,
) {

	flight := &domain.Flight{
		FlightID:    "AB123",
		Origin:      "JFK",
		Destination: "LAX",
		Departure:   time.Date(2024, 7, 10, 8, 0, 0, 0, time.Local),
		Seats: map[domain.SeatClass]*domain.SeatInfo{
			domain.NotAvailable: {Total: 1, Available: 1, BasePrice: 100, SeatMap: make(map[string]bool)}, // for test full or not available seat class
			domain.Economy:      {Total: 100, Available: 100, BasePrice: 300, SeatMap: make(map[string]bool)},
			domain.Business:     {Total: 30, Available: 30, BasePrice: 1000, SeatMap: make(map[string]bool)},
			domain.First:        {Total: 10, Available: 10, BasePrice: 3000, SeatMap: make(map[string]bool)},
		},
		ReservedSeats: map[domain.SeatClass]map[string]bool{
			domain.NotAvailable: make(map[string]bool), // for test full or not available seat class
			domain.Economy:      make(map[string]bool),
			domain.Business:     make(map[string]bool),
			domain.First:        make(map[string]bool),
		},
	}
	flightRepo.AddFlight(flight)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Book a flight")
		fmt.Println("2. Cancel flight")
		fmt.Println("3. Get flight information")
		fmt.Println("4. Search available flight based on origin, destination, and date")
		fmt.Println("5. Get Passenger detail")
		fmt.Println("9. Exit")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("invalid choice.")
			return
		}

		fmt.Println("====================================================================================")

		switch choice {
		case 1:
			BookFlight(bookingService)
		case 2:
			CancelFlight(bookingService)
		case 3:
			GetFlightInformation(flightService)
		case 4:
			SearchAvailableFlightBasedOn(flightRepo)
		case 5:
			GetPassengerDetails(passengerService)
		default:
			{
				fmt.Println("Exiting program.")
				return
			}

		}
	}
}

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
		if err != fmt.Errorf("no seats available in %s", seatClass) {
			fmt.Println("Booking failed: ", err)
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

		fmt.Printf("Booking failed: %v Upgrade to %s? (yes/no)\n", err, domain.ClassOrder[classIndex+1])
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

func SearchAvailableFlightBasedOn(flightRepo *flightRepo.FlightRepository) {

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

func GetPassengerDetails(passengerService *passengerSvc.PassengerService) {

	defer fmt.Println("====================================================================================")

	var passengerID string

	fmt.Println("Enter passenger id:")
	_, err := fmt.Scan(&passengerID)
	if err != nil {
		fmt.Println("invalid passenger id format.")
		return
	}

	passenger, err := passengerService.GetPassengerDetails(passengerID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		p, _ := json.Marshal(passenger)
		fmt.Printf("Passenger Details: %+v\n", string(p))
	}
}
