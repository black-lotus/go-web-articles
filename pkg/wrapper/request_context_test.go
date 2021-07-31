package wrapper

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/labstack/echo"
)

var (
	// request *RequestContext
	ctx echo.Context
)

func TestNewRequestContext(t *testing.T) {

	var storeID string = gofakeit.Word()
	var channelID string = gofakeit.Word()

	var header http.Header = http.Header{}
	header.Add("X-store-ID", storeID)
	header.Add("channelId", channelID)

	tests := []struct {
		name string
		want *RequestContext
	}{
		{
			name: "#test_1, positif case",
			want: &RequestContext{
				StoreID:   &storeID,
				ChannelID: &channelID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			request := &http.Request{
				Header: header,
			}

			context := e.NewContext(request, nil)

			result := NewRequestContext(context)

			if got := result; !reflect.DeepEqual(got.StoreID, tt.want.StoreID) {
				t.Errorf("NewRequestContext() = %v, want %v", got.StoreID, tt.want.StoreID)
			}
			if got := result; !reflect.DeepEqual(got.ChannelID, tt.want.ChannelID) {
				t.Errorf("NewRequestContext() = %v, want %v", got.ChannelID, tt.want.ChannelID)
			}
		})
	}
}

func Test_getHeader(t *testing.T) {
	var header http.Header = http.Header{}
	header.Add("x-store-id", "TIKETCOM")
	header.Add("channelId", "channel123")

	type args struct {
		header http.Header
		newKey string
		oldKey string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantError bool
	}{
		{
			name: "Positif case get from newKey",
			args: args{
				header: header,
				newKey: "x-store-id",
				oldKey: "",
			},
			want:      "TIKETCOM",
			wantError: false,
		},
		{
			name: "Positif case get from old key",
			args: args{
				header: header,
				newKey: "",
				oldKey: "channelId",
			},
			want:      "channel123",
			wantError: false,
		},
		{
			name: "Negative case",
			args: args{
				header: header,
				newKey: "asasas",
				oldKey: "vvvv",
			},
			want:      "TIKETCOM",
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHeader(tt.args.header, tt.args.newKey, tt.args.oldKey); got != tt.want {
				if !tt.wantError {
					t.Errorf("getHeader() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
