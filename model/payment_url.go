package govnpaymodels

import "time"

type GetPaymentURLRequest struct {
	Version        string
	TmnCode        string
	ReturnURL      string
	Amount         int64
	OrderInfo      string
	TxnRef         string
	CurrentCode    string
	OrderType      string
	CreateDate     time.Time
	TTL            time.Duration
	Locale         string
	IpAddr         string
	HashSecret     string
	HashAlgo       string
	InitPaymentURL string
}

func (req *GetPaymentURLRequest) GetVersion() string {
	if req != nil {
		return req.Version
	}

	return ""
}

func (req *GetPaymentURLRequest) GetTmnCode() string {
	if req != nil {
		return req.TmnCode
	}

	return ""
}

func (req *GetPaymentURLRequest) GetReturnURL() string {
	if req != nil {
		return req.ReturnURL
	}

	return ""
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

func (req *GetPaymentURLRequest) GetHashSecret() string {
	if req != nil {
		return req.HashSecret
	}
	return ""
}

func (req *GetPaymentURLRequest) GetHashAlgo() string {
	if req != nil {
		return req.HashAlgo
	}
	return ""
}

func (req *GetPaymentURLRequest) GetInitPaymentURL() string {
	if req != nil {
		return req.InitPaymentURL
	}
	return ""
}
