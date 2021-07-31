package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"reflect"
	"strconv"
	"time"

	my_error "webarticles/pkg/error"
)

const (
	// Success for success message
	Success = "SUCCESS"
	// Failed for failed message
	Failed = "FAILED"
)

type (
	Time time.Time
)

// HTTPResponse format
type HTTPResponse struct {
	ResponseCode int         `json:"-"`
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	Errors       interface{} `json:"errors,omitempty"`
	ServerTime   Time        `json:"serverTime"`
}

// NewHTTPResponse for create common response
func NewHTTPResponse(code int, message string, params ...interface{}) *HTTPResponse {
	commonResponse := new(HTTPResponse)

	for _, param := range params {
		switch e := param.(type) {
		case *my_error.MultiError:
		case error:
			param = my_error.NewMultiError().Append("detail", e)
		}

		// get value param if type is pointer
		refValue := reflect.ValueOf(param)
		if refValue.Kind() == reflect.Ptr {
			refValue = refValue.Elem()
		}
		param = refValue.Interface()

		switch val := param.(type) {
		case my_error.MultiError:
			commonResponse.Errors = val.ToMap()
		default:
			commonResponse.Data = param
		}
	}

	commonResponse.ServerTime = Time(time.Now())
	commonResponse.ResponseCode = code

	if code < http.StatusBadRequest {
		commonResponse.Code = Success
		commonResponse.Message = Success
		return commonResponse
	}

	commonResponse.Code = Failed
	commonResponse.Message = message

	return commonResponse
}

// JSON for set http JSON response (Content-Type: application/json) with parameter is http response writer
func (resp *HTTPResponse) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.ResponseCode)
	return json.NewEncoder(w).Encode(resp)
}

// XML for set http XML response (Content-Type: application/xml)
func (resp *HTTPResponse) XML(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(resp.ResponseCode)
	return xml.NewEncoder(w).Encode(resp)
}

// MarshalJSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON
func (t *Time) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

// Unix convert time to unix
func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

// Time convert time to UTC
func (t Time) Time() time.Time {
	return time.Time(t).UTC()
}

// String convert time to string format
func (t Time) String() string {
	return t.Time().String()
}
