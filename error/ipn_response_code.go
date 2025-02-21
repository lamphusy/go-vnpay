package govnpayerrors

import (
	"fmt"
)

type IPNResponseCode string

const (
	IPNCodeTransactionSuccess       IPNResponseCode = "00"
	IPNCodeSuspectedFraud           IPNResponseCode = "07"
	IPNCodeUnregisteredInternetBank IPNResponseCode = "09"
	IPNCodeInvalidCardInfo          IPNResponseCode = "10"
	IPNCodePaymentTimeout           IPNResponseCode = "11"
	IPNCodeCardLocked               IPNResponseCode = "12"
	IPNCodeInvalidOTP               IPNResponseCode = "13"
	IPNCodeTransactionCancelled     IPNResponseCode = "24"
	IPNCodeInsufficientFunds        IPNResponseCode = "51"
	IPNCodeExceededTransactionLimit IPNResponseCode = "65"
	IPNCodeBankMaintenance          IPNResponseCode = "75"
	IPNCodeExceededPasswordAttempts IPNResponseCode = "79"
	IPNCodeOtherErrors              IPNResponseCode = "99"
)

// ToString returns the response code as a two-digit string
func (code IPNResponseCode) ToString() string {
	return string(code)
}

// Message returns a detailed message for the response code
func (code IPNResponseCode) Message() string {
	switch code {
	case IPNCodeTransactionSuccess:
		return "Transaction successful."
	case IPNCodeSuspectedFraud:
		return "Transaction successful but suspected of fraud or unusual activity."
	case IPNCodeUnregisteredInternetBank:
		return "Transaction failed: The customer's card/account is not registered for Internet Banking."
	case IPNCodeInvalidCardInfo:
		return "Transaction failed: Customer entered incorrect card/account information more than 3 times."
	case IPNCodePaymentTimeout:
		return "Transaction failed: Payment timeout expired. Please try again."
	case IPNCodeCardLocked:
		return "Transaction failed: The customer's card/account is locked."
	case IPNCodeInvalidOTP:
		return "Transaction failed: Incorrect OTP entered. Please try again."
	case IPNCodeTransactionCancelled:
		return "Transaction failed: Customer canceled the transaction."
	case IPNCodeInsufficientFunds:
		return "Transaction failed: Insufficient funds in the account."
	case IPNCodeExceededTransactionLimit:
		return "Transaction failed: Daily transaction limit exceeded."
	case IPNCodeBankMaintenance:
		return "Transaction failed: The bank's payment system is under maintenance."
	case IPNCodeExceededPasswordAttempts:
		return "Transaction failed: Too many incorrect payment password attempts. Please try again."
	case IPNCodeOtherErrors:
		return "Transaction failed: Other errors (not listed in the predefined codes)."
	default:
		return fmt.Sprint("Unknown response code: ", code)
	}
}

func (code IPNResponseCode) IsIPNTransactionSuccess() bool {
	return code == IPNCodeTransactionSuccess
}

func (code IPNResponseCode) IsIPNSuspectedFraud() bool {
	return code == IPNCodeSuspectedFraud
}

func (code IPNResponseCode) IsIPNUnregisteredInternetBank() bool {
	return code == IPNCodeUnregisteredInternetBank
}

func (code IPNResponseCode) IsIPNInvalidCardInfo() bool {
	return code == IPNCodeInvalidCardInfo
}

func (code IPNResponseCode) IsIPNPaymentTimeout() bool {
	return code == IPNCodePaymentTimeout
}

func (code IPNResponseCode) IsIPNCardLocked() bool {
	return code == IPNCodeCardLocked
}

func (code IPNResponseCode) IsIPNInvalidOTP() bool {
	return code == IPNCodeInvalidOTP
}

func (code IPNResponseCode) IsIPNTransactionCancelled() bool {
	return code == IPNCodeTransactionCancelled
}

func (code IPNResponseCode) IsIPNInsufficientFunds() bool {
	return code == IPNCodeInsufficientFunds
}

func (code IPNResponseCode) IsIPNExceededTransactionLimit() bool {
	return code == IPNCodeExceededTransactionLimit
}

func (code IPNResponseCode) IsIPNBankMaintenance() bool {
	return code == IPNCodeBankMaintenance
}

func (code IPNResponseCode) IsIPNExceededPasswordAttempts() bool {
	return code == IPNCodeExceededPasswordAttempts
}

func (code IPNResponseCode) IsIPNOtherErrors() bool {
	return code == IPNCodeOtherErrors
}
