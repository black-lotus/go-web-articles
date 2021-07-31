package wrapper

import (
	"net/http"

	"github.com/labstack/echo"
)

// RequestContext for common HTTP header format
type RequestContext struct {
	StoreID        *string `validate:"required"`
	ChannelID      *string `validate:"required"`
	RequestID      *string `validate:"required"`
	ServiceID      *string `validate:"required"`
	Username       *string `validate:"required"`
	AcceptLanguage *string `validate:"required"`
	Currency       *string
	ResellerID     *string
	Identity       *string
	AccountID      *string
	BusinessID     *string
	LoginMedia     *string
	ForwardedFor   *string
	TrueClientIP   *string
}

const (
	// v3.0.0 format
	xStoreID        string = "X-Store-Id"
	xChannelID      string = "X-Channel-Id"
	xRequestID      string = "X-Request-Id"
	xServiceID      string = "X-Service-Id"
	xUsername       string = "X-Username"
	xAcceptLanguage string = "Accept-Language"
	xCurrenncy      string = "X-Currency"
	xResellerID     string = "X-Reseller-Id"
	xIdentity       string = "X-Identity"
	xAccountID      string = "X-Account-Id"
	xBusinessID     string = "X-Business-Id"
	xLoginMedia     string = "X-Login-Media"
	xForwardedFor   string = "X-Forwarded-For"
	xTrueClientIP   string = "True-Client-Ip"

	// old format
	oldStoreID    string = "storeId"
	oldChannelID  string = "channelId"
	oldRequestID  string = "requestId"
	oldServiceID  string = "serviceId"
	oldResellerID string = "resellerId"
	oldIdentity   string = "identity"
)

// NewRequestContext for create common http request header
func NewRequestContext(ctx echo.Context) *RequestContext {
	header := ctx.Request().Header
	var (
		storeID        string
		channelID      string
		requestID      string
		serviceID      string
		username       string
		acceptLanguage string
		currency       string
		resellerID     string
		identity       string
		accountID      string
		businessID     string
		loginMedia     string
		forwardedFor   string
		trueClientIP   string
	)

	storeID = getHeader(header, xStoreID, oldStoreID)
	channelID = getHeader(header, xChannelID, oldChannelID)
	requestID = getHeader(header, xRequestID, oldRequestID)
	serviceID = getHeader(header, xServiceID, oldServiceID)
	resellerID = getHeader(header, xResellerID, oldResellerID)
	identity = getHeader(header, xIdentity, oldIdentity)
	username = header.Get(xUsername)
	acceptLanguage = header.Get(acceptLanguage)
	currency = header.Get(xCurrenncy)
	accountID = header.Get(xAccountID)
	businessID = header.Get(xBusinessID)
	loginMedia = header.Get(xLoginMedia)
	forwardedFor = header.Get(xForwardedFor)
	trueClientIP = header.Get(xTrueClientIP)

	return &RequestContext{
		StoreID:        &storeID,
		ChannelID:      &channelID,
		RequestID:      &requestID,
		ServiceID:      &serviceID,
		Username:       &username,
		AcceptLanguage: &acceptLanguage,
		Currency:       &currency,
		ResellerID:     &resellerID,
		Identity:       &identity,
		AccountID:      &accountID,
		BusinessID:     &businessID,
		LoginMedia:     &loginMedia,
		ForwardedFor:   &forwardedFor,
		TrueClientIP:   &trueClientIP,
	}
}

func getHeader(header http.Header, newKey, oldKey string) string {
	var result string
	result = header.Get(newKey)
	if result == "" {
		result = header.Get(oldKey)
	}

	return result
}
