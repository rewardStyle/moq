// Code generated by moq; DO NOT EDIT.
// github.com/rewardStyle/moq

package genparamname

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

// Ensure, that InterfaceMock does implement Interface.
// If this is not the case, regenerate this file with moq.
var _ Interface = &InterfaceMock{}

// InterfaceMock is a mock implementation of Interface.
//
// 	func TestSomethingThatUsesInterface(t *testing.T) {
//
// 		// make and configure a mocked Interface
// 		mockedInterface := &InterfaceMock{
// 			MethodFunc: func(myTypeMoqParam *myType, numbers [3]json.Number, bytes []byte, nullStringToReader map[sql.NullString]io.Reader, fn func(conn net.Conn), goMoqParam Go, bufferPoolCh chan *httputil.BufferPool, val struct{URL *url.URL}, ifaceVal interface{CookieJar() http.CookieJar; fmt.Stringer})  {
// 				panic("mock out the Method method")
// 			},
// 		}
//
// 		// use mockedInterface in code that requires Interface
// 		// and then make assertions.
//
// 	}
type InterfaceMock struct {
	// MethodFunc mocks the Method method.
	MethodFunc func(myTypeMoqParam *myType, numbers [3]json.Number, bytes []byte, nullStringToReader map[sql.NullString]io.Reader, fn func(conn net.Conn), goMoqParam Go, bufferPoolCh chan *httputil.BufferPool, val struct{ URL *url.URL }, ifaceVal interface {
		CookieJar() http.CookieJar
		fmt.Stringer
	})

	// calls tracks calls to the methods.
	calls struct {
		// Method holds details about calls to the Method method.
		Method []struct {
			// MyTypeMoqParam is the myTypeMoqParam argument value.
			MyTypeMoqParam *myType
			// Numbers is the numbers argument value.
			Numbers [3]json.Number
			// Bytes is the bytes argument value.
			Bytes []byte
			// NullStringToReader is the nullStringToReader argument value.
			NullStringToReader map[sql.NullString]io.Reader
			// Fn is the fn argument value.
			Fn func(conn net.Conn)
			// GoMoqParam is the goMoqParam argument value.
			GoMoqParam Go
			// BufferPoolCh is the bufferPoolCh argument value.
			BufferPoolCh chan *httputil.BufferPool
			// Val is the val argument value.
			Val struct{ URL *url.URL }
			// IfaceVal is the ifaceVal argument value.
			IfaceVal interface {
				CookieJar() http.CookieJar
				fmt.Stringer
			}
		}
	}
	lockMethod sync.RWMutex
}

// Method calls MethodFunc.
func (mock *InterfaceMock) Method(myTypeMoqParam *myType, numbers [3]json.Number, bytes []byte, nullStringToReader map[sql.NullString]io.Reader, fn func(conn net.Conn), goMoqParam Go, bufferPoolCh chan *httputil.BufferPool, val struct{ URL *url.URL }, ifaceVal interface {
	CookieJar() http.CookieJar
	fmt.Stringer
}) {
	if mock.MethodFunc == nil {
		panic("InterfaceMock.MethodFunc: method is nil but Interface.Method was just called")
	}
	callInfo := struct {
		MyTypeMoqParam     *myType
		Numbers            [3]json.Number
		Bytes              []byte
		NullStringToReader map[sql.NullString]io.Reader
		Fn                 func(conn net.Conn)
		GoMoqParam         Go
		BufferPoolCh       chan *httputil.BufferPool
		Val                struct{ URL *url.URL }
		IfaceVal           interface {
			CookieJar() http.CookieJar
			fmt.Stringer
		}
	}{
		MyTypeMoqParam:     myTypeMoqParam,
		Numbers:            numbers,
		Bytes:              bytes,
		NullStringToReader: nullStringToReader,
		Fn:                 fn,
		GoMoqParam:         goMoqParam,
		BufferPoolCh:       bufferPoolCh,
		Val:                val,
		IfaceVal:           ifaceVal,
	}
	mock.lockMethod.Lock()
	mock.calls.Method = append(mock.calls.Method, callInfo)
	mock.lockMethod.Unlock()
	mock.MethodFunc(myTypeMoqParam, numbers, bytes, nullStringToReader, fn, goMoqParam, bufferPoolCh, val, ifaceVal)
}

// MethodCalls gets all the calls that were made to Method.
// Check the length with:
//     len(mockedInterface.MethodCalls())
func (mock *InterfaceMock) MethodCalls() []struct {
	MyTypeMoqParam     *myType
	Numbers            [3]json.Number
	Bytes              []byte
	NullStringToReader map[sql.NullString]io.Reader
	Fn                 func(conn net.Conn)
	GoMoqParam         Go
	BufferPoolCh       chan *httputil.BufferPool
	Val                struct{ URL *url.URL }
	IfaceVal           interface {
		CookieJar() http.CookieJar
		fmt.Stringer
	}
} {
	var calls []struct {
		MyTypeMoqParam     *myType
		Numbers            [3]json.Number
		Bytes              []byte
		NullStringToReader map[sql.NullString]io.Reader
		Fn                 func(conn net.Conn)
		GoMoqParam         Go
		BufferPoolCh       chan *httputil.BufferPool
		Val                struct{ URL *url.URL }
		IfaceVal           interface {
			CookieJar() http.CookieJar
			fmt.Stringer
		}
	}
	mock.lockMethod.RLock()
	calls = mock.calls.Method
	mock.lockMethod.RUnlock()
	return calls
}
