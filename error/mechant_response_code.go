package govnpayerrors

type MerchantResponseCode string

const (
	MerchantRespSuccess          MerchantResponseCode = "00"
	MerchantRespNotFound         MerchantResponseCode = "01"
	MerchantRespAlreadyConfirm   MerchantResponseCode = "02"
	MerchantRespInvalidAmount    MerchantResponseCode = "04"
	MerchantRespInvalidSignature MerchantResponseCode = "97"
	MerchantRespUnknowError      MerchantResponseCode = "99"
)

// ToString returns the response code as a two-digit string
func (code MerchantResponseCode) ToString() string {
	return string(code)
}

// Message returns a detailed message for the response code
func (code MerchantResponseCode) Message() string {
	switch code {
	case MerchantRespSuccess:
		return "Confirm Success"
	case MerchantRespNotFound:
		return "Transaction not found"
	case MerchantRespAlreadyConfirm:
		return "Order already confirmed"
	case MerchantRespInvalidAmount:
		return "Invalid amount"
	case MerchantRespInvalidSignature:
		return "Invalid Signature"
	case MerchantRespUnknowError:
		return "Unknown error"
	default:
		return "Unknown response code"
	}
}

func (code MerchantResponseCode) IsSuccess() bool {
	if code == MerchantRespSuccess {
		return true
	}
	return false
}

func (code MerchantResponseCode) IsNotFound() bool {
	if code == MerchantRespUnknowError {
		return true
	}
	return false
}

func (code MerchantResponseCode) IsAlreadyConfirm() bool {
	if code == MerchantRespAlreadyConfirm {
		return true
	}
	return false
}

func (code MerchantResponseCode) IsInvalidAmount() bool {
	if code == MerchantRespInvalidAmount {
		return true
	}
	return false
}

func (code MerchantResponseCode) IsInvalidSignature() bool {
	if code == MerchantRespInvalidSignature {
		return true
	}
	return false
}

func (code MerchantResponseCode) IsUnknowError() bool {
	if code == MerchantRespUnknowError {
		return true
	}
	return false
}
