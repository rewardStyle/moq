// Code generated by moq; DO NOT EDIT.
// github.com/rewardStyle/moq

package channels

import (
	"sync"
)

// Ensure, that QueuerMock does implement Queuer.
// If this is not the case, regenerate this file with moq.
var _ Queuer = &QueuerMock{}

// QueuerMock is a mock implementation of Queuer.
//
// 	func TestSomethingThatUsesQueuer(t *testing.T) {
//
// 		// make and configure a mocked Queuer
// 		mockedQueuer := &QueuerMock{
// 			SubFunc: func(topic string) (<-chan Queue, error) {
// 				panic("mock out the Sub method")
// 			},
// 			UnsubFunc: func(topic string)  {
// 				panic("mock out the Unsub method")
// 			},
// 		}
//
// 		// use mockedQueuer in code that requires Queuer
// 		// and then make assertions.
//
// 	}
type QueuerMock struct {
	// SubFunc mocks the Sub method.
	SubFunc func(topic string) (<-chan Queue, error)

	// UnsubFunc mocks the Unsub method.
	UnsubFunc func(topic string)

	// calls tracks calls to the methods.
	calls struct {
		// Sub holds details about calls to the Sub method.
		Sub []struct {
			// Topic is the topic argument value.
			Topic string
		}
		// Unsub holds details about calls to the Unsub method.
		Unsub []struct {
			// Topic is the topic argument value.
			Topic string
		}
	}
	lockSub   sync.RWMutex
	lockUnsub sync.RWMutex
}

// Sub calls SubFunc.
func (mock *QueuerMock) Sub(topic string) (<-chan Queue, error) {
	callInfo := struct {
		Topic string
	}{
		Topic: topic,
	}
	mock.lockSub.Lock()
	mock.calls.Sub = append(mock.calls.Sub, callInfo)
	mock.lockSub.Unlock()
	if mock.SubFunc == nil {
		var (
			queueChOut <-chan Queue
			errOut     error
		)
		return queueChOut, errOut
	}
	return mock.SubFunc(topic)
}

// SubCalls gets all the calls that were made to Sub.
// Check the length with:
//     len(mockedQueuer.SubCalls())
func (mock *QueuerMock) SubCalls() []struct {
	Topic string
} {
	var calls []struct {
		Topic string
	}
	mock.lockSub.RLock()
	calls = mock.calls.Sub
	mock.lockSub.RUnlock()
	return calls
}

// Unsub calls UnsubFunc.
func (mock *QueuerMock) Unsub(topic string) {
	callInfo := struct {
		Topic string
	}{
		Topic: topic,
	}
	mock.lockUnsub.Lock()
	mock.calls.Unsub = append(mock.calls.Unsub, callInfo)
	mock.lockUnsub.Unlock()
	if mock.UnsubFunc == nil {
		return
	}
	mock.UnsubFunc(topic)
}

// UnsubCalls gets all the calls that were made to Unsub.
// Check the length with:
//     len(mockedQueuer.UnsubCalls())
func (mock *QueuerMock) UnsubCalls() []struct {
	Topic string
} {
	var calls []struct {
		Topic string
	}
	mock.lockUnsub.RLock()
	calls = mock.calls.Unsub
	mock.lockUnsub.RUnlock()
	return calls
}
