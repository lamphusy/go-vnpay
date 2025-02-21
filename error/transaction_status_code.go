package vnpintgerrors

import (
	"fmt"
)

type TransactionStatus string

const (
	TransactionSuccess           TransactionStatus = "00"
	TransactionIncomplete        TransactionStatus = "01"
	TransactionError             TransactionStatus = "02"
	TransactionReversed          TransactionStatus = "04"
	TransactionRefundProcessing  TransactionStatus = "05"
	TransactionRefundRequestSent TransactionStatus = "06"
	TransactionSuspectedFraud    TransactionStatus = "07"
	TransactionRefundRejected    TransactionStatus = "09"
)

// ToStatusCode returns the transaction status as a two-digit string
func (status TransactionStatus) ToStatusCode() string {
	return string(status)
}

// Message returns a detailed message for the transaction status
func (status TransactionStatus) Message() string {
	switch status {
	case TransactionSuccess:
		return "Transaction successful."
	case TransactionIncomplete:
		return "Transaction incomplete."
	case TransactionError:
		return "Transaction encountered an error."
	case TransactionReversed:
		return "Transaction reversed: Customer's bank account was debited, but the transaction was not successful at VNPAY."
	case TransactionRefundProcessing:
		return "Transaction under processing: VNPAY is processing this refund transaction."
	case TransactionRefundRequestSent:
		return "Transaction refund request sent: VNPAY has sent the refund request to the bank."
	case TransactionSuspectedFraud:
		return "Transaction suspected of fraud."
	case TransactionRefundRejected:
		return "Transaction refund request rejected."
	default:
		return fmt.Sprintf("Transaction unknown status: %s", status)
	}
}

func (status TransactionStatus) IsTransactionSuccess() bool {
	return status == TransactionSuccess
}

func (status TransactionStatus) IsTransactionIncomplete() bool {
	return status == TransactionIncomplete
}

func (status TransactionStatus) IsTransactionError() bool {
	return status == TransactionError
}

func (status TransactionStatus) IsTransactionReversed() bool {
	return status == TransactionReversed
}

func (status TransactionStatus) IsTransactionRefundProcessing() bool {
	return status == TransactionRefundProcessing
}

func (status TransactionStatus) IsTransactionRefundRequestSent() bool {
	return status == TransactionRefundRequestSent
}

func (status TransactionStatus) IsTransactionSuspectedFraud() bool {
	return status == TransactionSuspectedFraud
}

func (status TransactionStatus) IsTransactionRefundRejected() bool {
	return status == TransactionRefundRejected
}
