# Flight booking system #

### This Go-based flight booking system supports: ###

- Seat reservation and cancellation
- Dynamic pricing
- Frequent flyer discounts
- Class upgrade options
- Show flight information
- Search available flights based on origin, destination, and departure date
- Passenger booking information

### Project Structure: ###

```plaintext
flight_booking_system/
├── cmd/                        
│ ├── root.go                       # entry point
│ ├── book_flight.go                # function book flight for interactive program
│ ├── cancel_flight.go              # function cancel flight for interactive program
│ ├── get_flight_info.go            # function get flight info for interactive program
│ ├── get_passenger_details.go      # function get passenger details for interactive program
│ ├── search_available_flight.go    # function search available flight for interactive program
│
├── constant/                       # constant variable
│ ├── constant.go
│ ├── err_msg.go
│
├── domain/                         # data struct
│ ├── booking.go
│ ├── flight.go
│ ├── passenger.go  
│
├── mocks/                          # mock files for test (auto-generated)
│
├── repository/                     # stored data
│ ├── flight
│ ├── ├── flight.go
│ ├── passenger
│ ├── ├── passenger.go
│
├── service/                        # main function source code
│ ├── booking
│ ├── ├── booking.go
│ ├── ├── booking_test.go
│ ├── ├── cancel_booking_test.go    # booking test
│ ├── flight
│ ├── ├── flight.go
│ ├── ├── flight_test.go            # flight test
│ ├── passenger
│ ├── ├── passenger.go
│ ├── ├── passenger_test.go         # passenger_test
│
├── .gitignore # Files to be ignored in Git
├── main.go
├── README.md # Instructions and documentation
```
### Installation guide ###

1. run Makefile: "make prepare" to install dependencies
2. to start the program, run Makefile: "make flight-book-system" or command "go run main.go flight-book-system"
3. to run test, run Makefile: "make test-coverage"
