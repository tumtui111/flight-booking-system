package domain

type Booking struct {
	BookingID    string    `json:"booking_id"`
	PassengerID  string    `json:"passenger_id"`
	FlightID     string    `json:"flight_id"`
	Seat         string    `json:"seat"`
	Price        float64   `json:"price"`
	Class        SeatClass `json:"class"`
	Status       string    `json:"status"`
	RefundAmount float64   `json:"refund_amount,omitempty"`
}
