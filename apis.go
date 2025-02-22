package govnpay

import (
	"fmt"
	"net/url"
	govnpayerrors "personal/vnpay-payment/error"
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
		return govnpayerrors.FormatInvalidArgumentError("Amount must be greater than zero")
	}

	if req.GetOrderInfo() == "" {
		return vnpintgerrors.FormatInvalidArgumentError("Order info is required")
	}

	if req.GetTxnRef() == "" {
		return vnpintgerrors.FormatInvalidArgumentError("Transaction reference is required")
	}

	if req.GetCurrentCode() == "" {
		return vnpintgerrors.FormatInvalidArgumentError("Current code is required")
	}

	if req.GetOrderType() == "" {
		return vnpintgerrors.FormatInvalidArgumentError("Order type is required")
	}

	if req.GetCreateDate().IsZero() {
		return vnpintgerrors.FormatInvalidArgumentError("Create date is required")
	}

	if req.GetTTL() == 0 {
		return vnpintgerrors.FormatInvalidArgumentError("Expire date is required")
	}

	if time.Now().Add(req.GetTTL()).Before(time.Now()) {
		return vnpintgerrors.FormatInvalidArgumentError("Time To Live (ttl) must be in the future")
	}

	if req.GetCreateDate().Add(req.GetTTL()).Before(req.GetCreateDate()) {
		return vnpintgerrors.FormatInvalidArgumentError("Time To Live (ttl) be after create date")
	}

	if req.GetLocale() == "" {
		return vnpintgerrors.FormatInvalidArgumentError("Locale is required")
	}

	if req.GetIpAddr() == "" {
		return vnpintgerrors.FormatInvalidArgumentError("IP address is required")
	}

	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		return "", vnpintgerrors.FormatInternalError("Cannot load time location: " + err.Error())
	}

	params := url.Values{}
	params.Add("vnp_Command", DefaultCommandPayment)
	params.Add("vnp_Amount", strconv.Itoa(int(r.GetAmount())*DefaultAmountFactor))
	params.Add("vnp_CreateDate", r.GetCreateDate().In(loc).Format(DefaultTimeFormat))
	params.Add("vnp_ExpireDate", time.Now().Add(r.GetTTL()).In(loc).Format(DefaultTimeFormat))

	params.Add("vnp_Version", hdl.GetConfig().GetVersion())
	params.Add("vnp_TmnCode", hdl.GetConfig().GetTmnCode())
	params.Add("vnp_ReturnUrl", hdl.GetConfig().GetReturnURL())

	params.Add("vnp_CurrCode", r.GetCurrentCode())
	params.Add("vnp_TxnRef", r.GetTxnRef())
	params.Add("vnp_OrderInfo", r.GetOrderInfo())
	params.Add("vnp_OrderType", r.GetOrderType())
	params.Add("vnp_Locale", r.GetLocale())
	params.Add("vnp_IpAddr", r.GetIpAddr())

	encodedParams := params.Encode()

	secureHash := hdl.computeSecureHash(encodedParams)
	encodedParams += "&vnp_SecureHash=" + secureHash

	respURL := hdl.GetConfig().GetInitPaymentURL() + "?" + encodedParams

	return respURL, nil
}
