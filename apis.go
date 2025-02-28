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

	if r.GetAmount() <= 0 {
		return "", fmt.Errorf("amount must be greater than zero")
	}

	if r.GetOrderInfo() == "" {
		return "", fmt.Errorf("order info is required")
	}

	if r.GetTxnRef() == "" {
		return "", fmt.Errorf("transaction reference is required")
	}

	if r.GetCurrentCode() == "" {
		return "", fmt.Errorf("current code is required")
	}

	if r.GetOrderType() == "" {
		return "", fmt.Errorf("order type is required")
	}

	if r.GetCreateDate().IsZero() {
		return "", fmt.Errorf("create date is required")
	}

	if r.GetTTL() == 0 {
		return "", fmt.Errorf("expire date is required")
	}

	if time.Now().Add(r.GetTTL()).Before(time.Now()) {
		return "", fmt.Errorf("time To Live (ttl) must be in the future")
	}

	if r.GetCreateDate().Add(r.GetTTL()).Before(r.GetCreateDate()) {
		return "", fmt.Errorf("time To Live (ttl) be after create date")
	}

	if r.GetLocale() == "" {
		return "", fmt.Errorf("locale is required")
	}

	if r.GetIpAddr() == "" {
		return "", fmt.Errorf("ip address is required")
	}

	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		return "", fmt.Errorf("cannot load time location: " + err.Error())
	}

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

	encodedParams := params.Encode()

	secureHash := helper.ComputeSecureHash(encodedParams, r.GetHashAlgo(), r.GetHashSecret())
	encodedParams += "&vnp_SecureHash=" + secureHash

	respURL := r.GetInitPaymentURL() + "?" + encodedParams

	return respURL, nil
}
