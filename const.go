package govnpay

import "time"

const (
	DefaultTimeFormat              = "20060102150405"
	DefaultAmountFactor            = 100
	DefaultCommandPayment          = "pay"
	DefaultCommandQueryTransaction = "querydr"
	DefaultLocale                  = "vn"
	DefaultOrderType               = "other"
	DefaultCurrentCode             = "VND"
	DefaultTimeZone                = "Asia/Ho_Chi_Minh"
	Version200                     = "2.0.0"
	Version201                     = "2.0.1"
	Version210                     = "2.1.0"
	DefaultMessageQueryTrans       = "Query transaction for request"
	DefaultMessagePayment          = "Payment for order"
	DefaultTTLPaymentURL           = 15 * time.Minute
	DefaultPrefixQueryTransReq     = "query."
)
