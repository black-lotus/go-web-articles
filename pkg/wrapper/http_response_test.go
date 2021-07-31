package wrapper

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	my_error "webarticles/pkg/error"

	"bou.ke/monkey"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPResponse(t *testing.T) {
	type Data struct {
		ID string `json:"id"`
	}

	multiError := my_error.NewMultiError()
	multiError.Append("test", fmt.Errorf("error test"))

	dateTime := time.Date(2021, 6, 30, 20, 34, 58, 651387237, time.UTC)

	type args struct {
		responseCode int
		message      string
		params       []interface{}
	}
	tests := []struct {
		name     string
		args     args
		mockFunc func()
		want     *HTTPResponse
	}{
		{
			name: "Testcase #1: Response data detail",
			args: args{
				responseCode: http.StatusOK,
				message:      "Get detail data",
				params: []interface{}{
					Data{ID: "061499700032"},
				},
			},
			mockFunc: func() {
				monkey.Patch(time.Now, func() time.Time {
					return dateTime
				})
			},
			want: &HTTPResponse{
				ResponseCode: 200,
				Code:         Success,
				Message:      Success,
				Data:         Data{ID: "061499700032"},
				ServerTime:   Time(dateTime),
			},
		},
		{
			name: "Testcase #2: Response only message (without data)",
			args: args{
				responseCode: http.StatusOK,
				message:      "list data empty",
			},
			mockFunc: func() {
				monkey.Patch(time.Now, func() time.Time {
					return dateTime
				})
			},
			want: &HTTPResponse{
				ResponseCode: 200,
				Code:         Success,
				Message:      Success,
				ServerTime:   Time(dateTime),
			},
		},
		{
			name: "Testcase #3: Response failed (ex: Bad Request)",
			args: args{
				responseCode: http.StatusBadRequest,
				message:      "id cannot be empty",
				params:       []interface{}{multiError},
			},
			mockFunc: func() {
				monkey.Patch(time.Now, func() time.Time {
					return dateTime
				})
			},
			want: &HTTPResponse{
				ResponseCode: 400,
				Code:         Failed,
				Message:      "id cannot be empty",
				Errors:       map[string]string{"test": "error test"},
				ServerTime:   Time(dateTime),
			},
		},
		{
			name: "Testcase #4: Response failed (error detail)",
			args: args{
				responseCode: http.StatusBadRequest,
				message:      "Failed validate",
				params:       []interface{}{errors.New("error")},
			},
			mockFunc: func() {
				monkey.Patch(time.Now, func() time.Time {
					return dateTime
				})
			},
			want: &HTTPResponse{
				ResponseCode: 400,
				Code:         Failed,
				Message:      "Failed validate",
				Errors:       map[string]string{"detail": "error"},
				ServerTime:   Time(dateTime),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got := NewHTTPResponse(tt.args.responseCode, tt.args.message, tt.args.params...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\x1b[31;1mNewHTTPResponse() = %v, \nwant => %v\x1b[0m", got, tt.want)
			}
		})
	}
}

func TestHTTPResponse_JSON(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/testing", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	resp := NewHTTPResponse(200, "success")
	assert.NoError(t, resp.JSON(c.Response()))
}

func TestHTTPResponse_XML(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/testing", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	resp := NewHTTPResponse(200, "success")
	assert.NoError(t, resp.XML(c.Response()))
}
