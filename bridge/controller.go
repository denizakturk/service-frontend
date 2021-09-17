package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller Struct

func NewServiceResponseError(message string) *ServiceResponseError {
	return &ServiceResponseError{s: message}
}

type ServiceResponseError struct {
	s string
}

func (r *ServiceResponseError) Error() string {
	return r.s
}
func (r ServiceResponseError) MarshalJSON() ([]byte, error) {
	if nil != &r.s {
		return []byte(`"` + r.Error() + `"`), nil
	}
	return []byte(""), nil
}

func (r ServiceResponseError) SetError(err error) {
	if nil != err {
		r.s = err.Error()
	}
}

type ServiceResponse struct {
	Error *ServiceResponseError `json:"error,omitempty"`
	Data  interface{}           `json:"data,omitempty"`
}

func (r *ServiceResponse) WriteResponse(w http.ResponseWriter) {
	if nil != r.Error {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
	response, _ := json.Marshal(*r)
	fmt.Fprintln(w, string(response))
	r.Error = nil
	r.Data = nil
}
