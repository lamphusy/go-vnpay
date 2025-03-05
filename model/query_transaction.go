package govnpaymodels

import "time"

type QueryTransactionRequest struct {
	RequestId       string
	TxnRef          string
	IpAddr          string
	OrderInfo       string
	TransactionDate time.Time
	CreateDate      time.Time
	HashSecret      string
	HashAlgo        string
	TmnCode         string
	Version         string
	QueryTransURL   string
}

func (req *QueryTransactionRequest) GetRequestId() string {
	if req != nil {
		return req.RequestId
	}
	return ""
}

func (req *QueryTransactionRequest) GetTxnRef() string {
	if req != nil {
		return req.TxnRef
	}
	return ""
}

func (req *QueryTransactionRequest) GetIpAddr() string {
	if req != nil {
		return req.IpAddr
	}
	return ""
}

func (req *QueryTransactionRequest) GetOrderInfo() string {
	if req != nil {
		return req.OrderInfo
	}
	return ""
}

func (req *QueryTransactionRequest) GetTransactionDate() time.Time {
	if req != nil {
		return req.TransactionDate
	}
	return time.Time{}
}

func (req *QueryTransactionRequest) GetCreateDate() time.Time {
	if req != nil {
		return req.CreateDate
	}
	return time.Time{}
}

func (req *QueryTransactionRequest) GetHashSecret() string {
	if req != nil {
		return req.HashSecret
	}
	return ""
}

func (req *QueryTransactionRequest) GetHashAlgo() string {
	if req != nil {
		return req.HashAlgo
	}
	return ""
}

func (req *QueryTransactionRequest) GetVersion() string {
	if req != nil {
		return req.GetVersion()
	}
	return ""
}

func (req *QueryTransactionRequest) GetTmnCode() string {
	if req != nil {
		return req.TmnCode
	}
	return ""
}

func (req *QueryTransactionRequest) GetQueryTransURL() string {
	if req != nil {
		return req.QueryTransURL
	}
	return ""
}

type VnPayQueryRequest struct {
	RequestId       string `json:"vnp_RequestId"`
	Version         string `json:"vnp_Version"`
	Command         string `json:"vnp_Command"`
	TmnCode         string `json:"vnp_TmnCode"`
	TxnRef          string `json:"vnp_TxnRef"`
	OrderInfo       string `json:"vnp_OrderInfo"`
	TransactionDate int64  `json:"vnp_TransactionDate"`
	CreateDate      int64  `json:"vnp_CreateDate"`
	IpAddr          string `json:"vnp_IpAddr"`
	SecureHash      string `json:"vnp_SecureHash"`
}

func (req *VnPayQueryRequest) GetRequestId() string {
	if req != nil {
		return req.RequestId
	}
	return ""
}

func (req *VnPayQueryRequest) GetVersion() string {
	if req != nil {
		return req.Version
	}
	return ""
}

func (req *VnPayQueryRequest) GetCommand() string {
	if req != nil {
		return req.Command
	}
	return ""
}

func (req *VnPayQueryRequest) GetTmnCode() string {
	if req != nil {
		return req.TmnCode
	}
	return ""
}

func (req *VnPayQueryRequest) GetTxnRef() string {
	if req != nil {
		return req.TxnRef
	}
	return ""
}

func (req *VnPayQueryRequest) GetOrderInfo() string {
	if req != nil {
		return req.OrderInfo
	}
	return ""
}

func (req *VnPayQueryRequest) GetTransactionDate() int64 {
	if req != nil {
		return req.TransactionDate
	}
	return 0
}

func (req *VnPayQueryRequest) GetCreateDate() int64 {
	if req != nil {
		return req.CreateDate
	}
	return 0
}

func (req *VnPayQueryRequest) GetIpAddr() string {
	if req != nil {
		return req.IpAddr
	}
	return ""
}

func (req *VnPayQueryRequest) GetSecureHash() string {
	if req != nil {
		return req.SecureHash
	}
	return ""
}

type VnPayQueryResponse struct {
	ResponseId        string `json:"vnp_ResponseId"`
	Command           string `json:"vnp_Command"`
	TmnCode           string `json:"vnp_TmnCode"`
	TxnRef            string `json:"vnp_TxnRef"`
	Amount            string `json:"vnp_Amount"`
	OrderInfo         string `json:"vnp_OrderInfo"`
	ResponseCode      string `json:"vnp_ResponseCode"`
	Message           string `json:"vnp_Message"`
	BankCode          string `json:"vnp_BankCode"`
	PayDate           string `json:"vnp_PayDate"`
	TransactionNo     string `json:"vnp_TransactionNo"`
	TransactionType   string `json:"vnp_TransactionType"`
	TransactionStatus string `json:"vnp_TransactionStatus"`
	PromotionCode     string `json:"vnp_PromotionCode"`
	PromotionAmount   string `json:"vnp_PromotionAmount"`
	SecureHash        string `json:"vnp_SecureHash"`
}

func (res *VnPayQueryResponse) GetResponseId() string {
	if res != nil {
		return res.ResponseId
	}
	return ""
}

func (res *VnPayQueryResponse) GetCommand() string {
	if res != nil {
		return res.Command
	}
	return ""
}

func (res *VnPayQueryResponse) GetTmnCode() string {
	if res != nil {
		return res.TmnCode
	}
	return ""
}

func (res *VnPayQueryResponse) GetTxnRef() string {
	if res != nil {
		return res.TxnRef
	}
	return ""
}

func (res *VnPayQueryResponse) GetAmount() string {
	if res != nil {
		return res.Amount
	}
	return ""
}

func (res *VnPayQueryResponse) GetOrderInfo() string {
	if res != nil {
		return res.OrderInfo
	}
	return ""
}

func (res *VnPayQueryResponse) GetResponseCode() string {
	if res != nil {
		return res.ResponseCode
	}
	return ""
}

func (res *VnPayQueryResponse) GetMessage() string {
	if res != nil {
		return res.Message
	}
	return ""
}

func (res *VnPayQueryResponse) GetBankCode() string {
	if res != nil {
		return res.BankCode
	}
	return ""
}

func (res *VnPayQueryResponse) GetPayDate() string {
	if res != nil {
		return res.PayDate
	}
	return ""
}

func (res *VnPayQueryResponse) GetTransactionNo() string {
	if res != nil {
		return res.TransactionNo
	}
	return ""
}

func (res *VnPayQueryResponse) GetTransactionType() string {
	if res != nil {
		return res.TransactionType
	}
	return ""
}

func (res *VnPayQueryResponse) GetTransactionStatus() string {
	if res != nil {
		return res.TransactionStatus
	}
	return ""
}

func (res *VnPayQueryResponse) GetPromotionCode() string {
	if res != nil {
		return res.PromotionCode
	}
	return ""
}

func (res *VnPayQueryResponse) GetPromotionAmount() string {
	if res != nil {
		return res.PromotionAmount
	}
	return ""
}

func (res *VnPayQueryResponse) GetSecureHash() string {
	if res != nil {
		return res.SecureHash
	}
	return ""
}

type QueryTransactionResponse struct {
	ResponseId        string `json:"response_id"`
	Command           string `json:"command"`
	TmnCode           string `json:"tmn_code"`
	TxnRef            string `json:"txn_ref"`
	Amount            int64  `json:"amount"`
	OrderInfo         string `json:"order_info"`
	ResponseCode      string `json:"response_code"`
	Message           string `json:"message"`
	BankCode          string `json:"bank_code"`
	PayDate           int64  `json:"pay_date"`
	TransactionNo     int64  `json:"transaction_no"`
	TransactionType   int32  `json:"transaction_type"`
	TransactionStatus string `json:"transaction_status"`
	PromotionCode     int64  `json:"promotion_code"`
	PromotionAmount   int64  `json:"promotion_amount"`
	SecureHash        string `json:"secure_hash"`
}

func (res *QueryTransactionResponse) GetResponseId() string {
	if res != nil {
		return res.ResponseId
	}
	return ""
}

func (res *QueryTransactionResponse) GetCommand() string {
	if res != nil {
		return res.Command
	}
	return ""
}

func (res *QueryTransactionResponse) GetTmnCode() string {
	if res != nil {
		return res.TmnCode
	}
	return ""
}

func (res *QueryTransactionResponse) GetTxnRef() string {
	if res != nil {
		return res.TxnRef
	}
	return ""
}

func (res *QueryTransactionResponse) GetAmount() int64 {
	if res != nil {
		return res.Amount
	}
	return 0
}

func (res *QueryTransactionResponse) GetOrderInfo() string {
	if res != nil {
		return res.OrderInfo
	}
	return ""
}

func (res *QueryTransactionResponse) GetResponseCode() string {
	if res != nil {
		return res.ResponseCode
	}
	return ""
}

func (res *QueryTransactionResponse) GetMessage() string {
	if res != nil {
		return res.Message
	}
	return ""
}

func (res *QueryTransactionResponse) GetBankCode() string {
	if res != nil {
		return res.BankCode
	}
	return ""
}

func (res *QueryTransactionResponse) GetPayDate() int64 {
	if res != nil {
		return res.PayDate
	}
	return 0
}

func (res *QueryTransactionResponse) GetTransactionNo() int64 {
	if res != nil {
		return res.TransactionNo
	}
	return 0
}

func (res *QueryTransactionResponse) GetTransactionType() int32 {
	if res != nil {
		return res.TransactionType
	}
	return 0
}

func (res *QueryTransactionResponse) GetTransactionStatus() string {
	if res != nil {
		return res.TransactionStatus
	}
	return ""
}

func (res *QueryTransactionResponse) GetPromotionCode() int64 {
	if res != nil {
		return res.PromotionCode
	}
	return 0
}

func (res *QueryTransactionResponse) GetPromotionAmount() int64 {
	if res != nil {
		return res.PromotionAmount
	}
	return 0
}

func (res *QueryTransactionResponse) GetSecureHash() string {
	if res != nil {
		return res.SecureHash
	}
	return ""
}
