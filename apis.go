package govnpay

import (
	"fmt"
	"net/url"
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
		r.OrderInfo = fmt.Sprintf("%s: %s", DefaultMessagePayment, req.GetTxnRef())
	}

	if r.GetTTL() == 0 {
		r.TTL = DefaultTTLPaymentURL
	}

	err := hdl.validateGetPaymentURLruest(r)
	if err != nil {
		return "", err
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
