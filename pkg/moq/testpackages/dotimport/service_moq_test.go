// Code generated by moq; DO NOT EDIT.
// github.com/rewardStyle/moq

package dotimport_test

import (
	"github.com/rewardStyle/moq/pkg/moq/testpackages/dotimport"
	"sync"
)

// Ensure, that ServiceMock does implement dotimport.Service.
// If this is not the case, regenerate this file with moq.
var _ dotimport.Service = &ServiceMock{}

// ServiceMock is a mock implementation of dotimport.Service.
//
// 	func TestSomethingThatUsesService(t *testing.T) {
//
// 		// make and configure a mocked dotimport.Service
// 		mockedService := &ServiceMock{
// 			UserFunc: func(ID string) (dotimport.User, error) {
// 				panic("mock out the User method")
// 			},
// 		}
//
// 		// use mockedService in code that requires dotimport.Service
// 		// and then make assertions.
//
// 	}
type ServiceMock struct {
	// UserFunc mocks the User method.
	UserFunc func(ID string) (dotimport.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// User holds details about calls to the User method.
		User []struct {
			// ID is the ID argument value.
			ID string
		}
	}
	lockUser sync.RWMutex
}

// User calls UserFunc.
func (mock *ServiceMock) User(ID string) (dotimport.User, error) {
	if mock.UserFunc == nil {
		panic("ServiceMock.UserFunc: method is nil but Service.User was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	mock.lockUser.Lock()
	mock.calls.User = append(mock.calls.User, callInfo)
	mock.lockUser.Unlock()
	return mock.UserFunc(ID)
}

// UserCalls gets all the calls that were made to User.
// Check the length with:
//     len(mockedService.UserCalls())
func (mock *ServiceMock) UserCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockUser.RLock()
	calls = mock.calls.User
	mock.lockUser.RUnlock()
	return calls
}
