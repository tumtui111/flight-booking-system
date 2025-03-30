package cmd

import (
	"encoding/json"
	"fmt"

	passengerSvc "flight-book-system/service/passenger"
)

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
