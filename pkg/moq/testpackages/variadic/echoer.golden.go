// Code generated by moq; DO NOT EDIT.
// github.com/rewardStyle/moq

package variadic

import (
	"sync"
)

// Ensure, that EchoerMock does implement Echoer.
// If this is not the case, regenerate this file with moq.
var _ Echoer = &EchoerMock{}

// EchoerMock is a mock implementation of Echoer.
//
// 	func TestSomethingThatUsesEchoer(t *testing.T) {
//
// 		// make and configure a mocked Echoer
// 		mockedEchoer := &EchoerMock{
// 			EchoFunc: func(ss ...string) []string {
// 				panic("mock out the Echo method")
// 			},
// 		}
//
// 		// use mockedEchoer in code that requires Echoer
// 		// and then make assertions.
//
// 	}
type EchoerMock struct {
	// EchoFunc mocks the Echo method.
	EchoFunc func(ss ...string) []string

	// calls tracks calls to the methods.
	calls struct {
		// Echo holds details about calls to the Echo method.
		Echo []struct {
			// Ss is the ss argument value.
			Ss []string
		}
	}
	lockEcho sync.RWMutex
}

// Echo calls EchoFunc.
func (mock *EchoerMock) Echo(ss ...string) []string {
	if mock.EchoFunc == nil {
		panic("EchoerMock.EchoFunc: method is nil but Echoer.Echo was just called")
	}
	callInfo := struct {
		Ss []string
	}{
		Ss: ss,
	}
	mock.lockEcho.Lock()
	mock.calls.Echo = append(mock.calls.Echo, callInfo)
	mock.lockEcho.Unlock()
	return mock.EchoFunc(ss...)
}

// EchoCalls gets all the calls that were made to Echo.
// Check the length with:
//     len(mockedEchoer.EchoCalls())
func (mock *EchoerMock) EchoCalls() []struct {
	Ss []string
} {
	var calls []struct {
		Ss []string
	}
	mock.lockEcho.RLock()
	calls = mock.calls.Echo
	mock.lockEcho.RUnlock()
	return calls
}
