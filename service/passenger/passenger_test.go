package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"flight-book-system/domain"
	mockPassenger "flight-book-system/mocks/service/passenger"
)

func Test_Passenger(t *testing.T) {
	t.Run("Success case | GetPassengerDetail | found data", func(t *testing.T) {

		passengerDetail := domain.Passenger{
			PassengerID: "AB12345",
			IsFrequentFlyer: false,

		}

		var (
			mockPassenger = mockPassenger.NewIPassengerRepository(t)
			PassengerService = NewPassengerService(mockPassenger)
		)

		mockPassenger.On("GetPassenger", "AB12345").Return(&passengerDetail, true)

		resp, err := PassengerService.GetPassengerDetails("AB12345")
		assert.Nil(t, err)
		assert.Equal(t, resp.PassengerID, "AB12345")

	})

	t.Run("Fail case | GetPassengerDetail | not found data", func(t *testing.T) {

		var (
			mockPassenger = mockPassenger.NewIPassengerRepository(t)
			PassengerService = NewPassengerService(mockPassenger)
		)

		mockPassenger.On("GetPassenger", "AB12345").Return(nil, false)

		_, err := PassengerService.GetPassengerDetails("AB12345")
		assert.NotNil(t, err)

	})
}
