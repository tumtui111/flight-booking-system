// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	domain "flight-book-system/domain"

	mock "github.com/stretchr/testify/mock"
)

// IFlightRepository is an autogenerated mock type for the IFlightRepository type
type IFlightRepository struct {
	mock.Mock
}

// GetFlight provides a mock function with given fields: flightID
func (_m *IFlightRepository) GetFlight(flightID string) (*domain.Flight, bool) {
	ret := _m.Called(flightID)

	if len(ret) == 0 {
		panic("no return value specified for GetFlight")
	}

	var r0 *domain.Flight
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (*domain.Flight, bool)); ok {
		return rf(flightID)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Flight); ok {
		r0 = rf(flightID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Flight)
		}
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(flightID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewIFlightRepository creates a new instance of IFlightRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIFlightRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IFlightRepository {
	mock := &IFlightRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
