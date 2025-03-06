package govnpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lamphusy/go-vnpay/helper"
	"github.com/lamphusy/go-vnpay/model"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetPaymentURL(r *govnpaymodels.GetPaymentURLRequest) (string, error) {
	// Set default values
	if r == nil {
		return "", fmt.Errorf("request cannot be nil")
	}

	setPaymentURLDefaults(r)

	// Validate request
	if err := validatePaymentURLRequest(r); err != nil {
		return "", err
	}

	// Get location
	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		return "", fmt.Errorf("cannot load time location: %w", err)
	}

	// Build URL parameters
	params := buildPaymentURLParams(r, loc)
	encodedParams := params.Encode()

	// Compute and append secure hash
	secureHash := helper.ComputeSecureHash(encodedParams, r.GetHashAlgo(), r.GetHashSecret())
	encodedParams += "&vnp_SecureHash=" + secureHash

	return r.GetInitPaymentURL() + "?" + encodedParams, nil
}

func setPaymentURLDefaults(r *govnpaymodels.GetPaymentURLRequest) {
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

func validatePaymentURLRequest(r *govnpaymodels.GetPaymentURLRequest) error {
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

func buildPaymentURLParams(r *govnpaymodels.GetPaymentURLRequest, loc *time.Location) url.Values {
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

func QueryTransaction(ctx context.Context, r *govnpaymodels.QueryTransactionRequest) (*govnpaymodels.QueryTransactionResponse, error) {
	fillQueryTransactionDefaults(r)

	err := validateQueryTransaction(r)
	if err != nil {
		return nil, err
	}

	reqToVNPay, err := buildVNPayQueryRequest(r)
	if err != nil {
		return nil, err
	}

	buf, err := sendHTTPRequest(ctx, r.GetQueryTransURL(), reqToVNPay)
	if err != nil {
		return nil, fmt.Errorf("send http request error: %w", err)
	}

	resp := &govnpaymodels.VnPayQueryResponse{}
	if err = json.Unmarshal(buf, resp); err != nil {
		return nil, fmt.Errorf("cannot unmarshal query transaction response error " + err.Error())
	}

	secureHash := computeResponseHash(resp, r.GetHashAlgo(), r.GetHashSecret())
	if resp.GetSecureHash() != secureHash {
		return nil, fmt.Errorf("invalid secure hash")
	}

	respToReturn := convertVNPayToQueryResponse(resp)

	return respToReturn, nil
}

func fillQueryTransactionDefaults(r *govnpaymodels.QueryTransactionRequest) {
	if r.GetRequestId() == "" {
		r.RequestId = strings.Replace(uuid.NewString(), "-", "", -1)
	}
	if r.GetOrderInfo() == "" {
		r.OrderInfo = fmt.Sprintf("%s: %s", DefaultMessageQueryTrans, r.GetRequestId())
	}
	if r.GetCreateDate().IsZero() {
		r.CreateDate = time.Now()
	}
}

func validateQueryTransaction(r *govnpaymodels.QueryTransactionRequest) error {
	// Validate required string fields
	requiredFields := map[string]string{
		"Request ID":            r.GetRequestId(),
		"Transaction reference": r.GetTxnRef(),
		"Order info":            r.GetOrderInfo(),
		"IP address":            r.GetIpAddr(),
	}

	for field, value := range requiredFields {
		if value == "" {
			return fmt.Errorf("%s is required", field)
		}
	}

	// Validate required dates
	requiredDates := map[string]time.Time{
		"Transaction date": r.GetTransactionDate(),
		"Create date":      r.GetCreateDate(),
	}

	for field, value := range requiredDates {
		if value.IsZero() {
			return fmt.Errorf(field + " is required")
		}
	}

	// Validate date relationships
	if r.GetCreateDate().Before(r.GetTransactionDate()) {
		return fmt.Errorf("create date must be after transaction date")
	}

	return nil
}

func buildVNPayQueryRequest(r *govnpaymodels.QueryTransactionRequest) (*govnpaymodels.VnPayQueryRequest, error) {
	loc, err := time.LoadLocation(DefaultTimeZone)
	if err != nil {
		return nil, fmt.Errorf("cannot load time location: " + err.Error())
	}

	transDate, err := strconv.ParseInt(r.GetTransactionDate().In(loc).Format(DefaultTimeFormat), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse transaction date: " + err.Error())
	}

	createDate, err := strconv.ParseInt(r.GetCreateDate().In(loc).Format(DefaultTimeFormat), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse create date: " + err.Error())
	}

	resp := &govnpaymodels.VnPayQueryRequest{
		RequestId:       r.GetRequestId(),
		Version:         r.GetVersion(),
		Command:         DefaultCommandQueryTransaction,
		TmnCode:         r.GetTmnCode(),
		TxnRef:          r.GetTxnRef(),
		OrderInfo:       r.GetOrderInfo(),
		TransactionDate: transDate,
		CreateDate:      createDate,
		IpAddr:          r.GetIpAddr(),
	}

	resp.SecureHash = computeRequestHash(resp, r.GetHashAlgo(), r.GetHashSecret())

	return resp, nil
}

func sendHTTPRequest(ctx context.Context, url string, reqToVNPay interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(reqToVNPay)
	if err != nil {
		return nil, fmt.Errorf("marshal request error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func computeRequestHash(r *govnpaymodels.VnPayQueryRequest, hashAlgo string, hashSecret string) string {
	hashData := r.GetRequestId() + "|" +
		r.GetVersion() + "|" +
		DefaultCommandQueryTransaction + "|" +
		r.GetTmnCode() + "|" +
		r.GetTxnRef() + "|" +
		fmt.Sprintf("%d", r.GetTransactionDate()) + "|" +
		fmt.Sprintf("%d", r.GetCreateDate()) + "|" +
		r.GetIpAddr() + "|" +
		r.GetOrderInfo()

	return helper.ComputeSecureHash(hashData, hashAlgo, hashSecret)
}

func computeResponseHash(data *govnpaymodels.VnPayQueryResponse, hashAlgo string, hashSecret string) string {
	hashData := data.GetResponseId() + "|" +
		data.GetCommand() + "|" +
		data.GetResponseCode() + "|" +
		data.GetMessage() + "|" +
		data.GetTmnCode() + "|" +
		data.GetTxnRef() + "|" +
		data.GetAmount() + "|" +
		data.GetBankCode() + "|" +
		data.GetPayDate() + "|" +
		data.GetTransactionNo() + "|" +
		data.GetTransactionType() + "|" +
		data.GetTransactionStatus() + "|" +
		data.GetOrderInfo() + "|" +
		data.GetPromotionCode() + "|" +
		data.GetPromotionAmount()

	return helper.ComputeSecureHash(hashData, hashAlgo, hashSecret)
}

func convertVNPayToQueryResponse(vnpRes *govnpaymodels.VnPayQueryResponse) *govnpaymodels.QueryTransactionResponse {
	if vnpRes == nil {
		return nil
	}

	return &govnpaymodels.QueryTransactionResponse{
		ResponseId:        vnpRes.GetResponseId(),
		Command:           vnpRes.GetCommand(),
		TmnCode:           vnpRes.GetTmnCode(),
		TxnRef:            vnpRes.GetTxnRef(),
		Amount:            helper.ParseAmount(vnpRes.GetAmount()),
		OrderInfo:         vnpRes.GetOrderInfo(),
		ResponseCode:      vnpRes.GetResponseCode(),
		Message:           vnpRes.GetMessage(),
		BankCode:          vnpRes.GetBankCode(),
		PayDate:           helper.ParseInt64(vnpRes.GetPayDate()),
		TransactionNo:     helper.ParseInt64(vnpRes.GetTransactionNo()),
		TransactionType:   helper.ParseInt32(vnpRes.GetTransactionType()),
		TransactionStatus: vnpRes.GetTransactionStatus(),
		PromotionCode:     helper.ParseInt64(vnpRes.GetPromotionCode()),
		PromotionAmount:   helper.ParseInt64(vnpRes.GetPromotionAmount()),
		SecureHash:        vnpRes.GetSecureHash(),
	}
}
