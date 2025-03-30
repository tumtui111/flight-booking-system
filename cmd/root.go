/*
Copyright © 2025 worawit.l.official@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
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
		Departure:   time.Date(2025, 4, 10, 8, 0, 0, 0, time.Local),
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
			SearchAvailableFlight(flightRepo)
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
