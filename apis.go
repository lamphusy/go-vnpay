package govnpay

import (
	"fmt"
	"net/url"
	"personal/vnpay-payment/helper"
	"personal/vnpay-payment/model"
	"strconv"
	"time"
)

func GetPaymentURL(r *govnpaymodels.GetPaymentURLRequest) (string, error) {
	// Set default values
	if r == nil {
		return "", fmt.Errorf("request cannot be nil")
	}

	setDefaults(r)

	// Validate request
	if err := validateRequest(r); err != nil {
		return "", err
	}

	// Get location
	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		return "", fmt.Errorf("cannot load time location: %w", err)
	}

	// Build URL parameters
	params := buildURLParams(r, loc)
	encodedParams := params.Encode()

	// Compute and append secure hash
	secureHash := helper.ComputeSecureHash(encodedParams, r.GetHashAlgo(), r.GetHashSecret())
	encodedParams += "&vnp_SecureHash=" + secureHash

	return r.GetInitPaymentURL() + "?" + encodedParams, nil
}

func setDefaults(r *govnpaymodels.GetPaymentURLRequest) {
	if r.GetLocale() == "" {
		r.Locale = DefaultLocale
	}
	if r.GetCurrentCode() == "" {
		r.CurrentCode = DefaultCurrentCode
	}
	if r.GetOrderType() == "" {
		r.OrderType = DefaultOrderType
	}
	if r.GetCreateDate().IsZero() {
		r.CreateDate = time.Now()
	}
	if r.GetOrderInfo() == "" {
		r.OrderInfo = fmt.Sprintf("%s: %s", DefaultMessagePayment, r.GetTxnRef())
	}
	if r.GetTTL() == 0 {
		r.TTL = DefaultTTLPaymentURL
	}
}

func validateRequest(r *govnpaymodels.GetPaymentURLRequest) error {
	if r.GetAmount() <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	requiredFields := map[string]string{
		"order info":            r.GetOrderInfo(),
		"transaction reference": r.GetTxnRef(),
		"current code":          r.GetCurrentCode(),
		"order type":            r.GetOrderType(),
		"locale":                r.GetLocale(),
		"ip address":            r.GetIpAddr(),
	}

	for field, value := range requiredFields {
		if value == "" {
			return fmt.Errorf("%s is required", field)
		}
	}

	now := time.Now()
	if r.GetCreateDate().IsZero() {
		return fmt.Errorf("create date is required")
	}
	if r.GetTTL() == 0 {
		return fmt.Errorf("expire date is required")
	}
	if now.Add(r.GetTTL()).Before(now) {
		return fmt.Errorf("time To Live (ttl) must be in the future")
	}
	if r.GetCreateDate().Add(r.GetTTL()).Before(r.GetCreateDate()) {
		return fmt.Errorf("time To Live (ttl) be after create date")
	}

	return nil
}

func buildURLParams(r *govnpaymodels.GetPaymentURLRequest, loc *time.Location) url.Values {
	params := url.Values{}
	params.Add("vnp_Command", DefaultCommandPayment)
	params.Add("vnp_Amount", strconv.Itoa(int(r.GetAmount())*DefaultAmountFactor))
	params.Add("vnp_CreateDate", r.GetCreateDate().In(loc).Format(DefaultTimeFormat))
	params.Add("vnp_ExpireDate", time.Now().Add(r.GetTTL()).In(loc).Format(DefaultTimeFormat))
	params.Add("vnp_Version", r.GetVersion())
	params.Add("vnp_TmnCode", r.GetTmnCode())
	params.Add("vnp_ReturnUrl", r.GetReturnURL())
	params.Add("vnp_CurrCode", r.GetCurrentCode())
	params.Add("vnp_TxnRef", r.GetTxnRef())
	params.Add("vnp_OrderInfo", r.GetOrderInfo())
	params.Add("vnp_OrderType", r.GetOrderType())
	params.Add("vnp_Locale", r.GetLocale())
	params.Add("vnp_IpAddr", r.GetIpAddr())
	return params
}
