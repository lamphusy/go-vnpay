package govnpaymodels

import "time"

type GetPaymentURLRequest struct {
	Amount      int64
	OrderInfo   string
	TxnRef      string
	CurrentCode string
	OrderType   string
	CreateDate  time.Time
	TTL         time.Duration
	Locale      string
	IpAddr      string
}

func (req *GetPaymentURLRequest) GetAmount() int64 {
	if req != nil {
		return req.Amount
	}

	return 0
}

func (req *GetPaymentURLRequest) GetOrderInfo() string {
	if req != nil {
		return req.OrderInfo
	}
	return ""
}

func (req *GetPaymentURLRequest) GetTxnRef() string {
	if req != nil {
		return req.TxnRef
	}
	return ""
}

func (req *GetPaymentURLRequest) GetCurrentCode() string {
	if req != nil {
		return req.CurrentCode
	}
	return ""
}

func (req *GetPaymentURLRequest) GetOrderType() string {
	if req != nil {
		return req.OrderType
	}
	return ""
}

func (req *GetPaymentURLRequest) GetCreateDate() time.Time {
	if req != nil {
		return req.CreateDate
	}
	return time.Time{}
}

func (req *GetPaymentURLRequest) GetTTL() time.Duration {
	if req != nil {
		return req.TTL
	}

	return 0
}

func (req *GetPaymentURLRequest) GetLocale() string {
	if req != nil {
		return req.Locale
	}

	return ""
}

func (req *GetPaymentURLRequest) GetIpAddr() string {
	if req != nil {
		return req.IpAddr
	}
	return ""
}
