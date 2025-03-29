// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "flight-book-system/domain"

	mock "github.com/stretchr/testify/mock"
)

// IPassengerRepository is an autogenerated mock type for the IPassengerRepository type
type IPassengerRepository struct {
	mock.Mock
}

// AddPassenger provides a mock function with given fields: passenger
func (_m *IPassengerRepository) AddPassenger(passenger *domain.Passenger) {
	_m.Called(passenger)
}

// GetPassenger provides a mock function with given fields: passengerID
func (_m *IPassengerRepository) GetPassenger(passengerID string) (*domain.Passenger, bool) {
	ret := _m.Called(passengerID)

	if len(ret) == 0 {
		panic("no return value specified for GetPassenger")
	}

	var r0 *domain.Passenger
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (*domain.Passenger, bool)); ok {
		return rf(passengerID)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Passenger); ok {
		r0 = rf(passengerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Passenger)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(passengerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// UpdatePassengerBookingRefundAmount provides a mock function with given fields: booking, refundAmount
func (_m *IPassengerRepository) UpdatePassengerBookingRefundAmount(booking *domain.Booking, refundAmount float64) {
	_m.Called(booking, refundAmount)
}

// UpdatePassengerBookingStatus provides a mock function with given fields: booking, status
func (_m *IPassengerRepository) UpdatePassengerBookingStatus(booking *domain.Booking, status string) {
	_m.Called(booking, status)
}

// NewIPassengerRepository creates a new instance of IPassengerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPassengerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPassengerRepository {
	mock := &IPassengerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
