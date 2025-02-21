package govnpayerrors

import (
	"fmt"
)

type QueryResponseCode string

const (
	QueryCodeRequestSuccess      QueryResponseCode = "00"
	QueryCodeInvalidConnectionID QueryResponseCode = "02"
	QueryCodeInvalidDataFormat   QueryResponseCode = "03"
	QueryCodeTransactionNotFound QueryResponseCode = "91"
	QueryCodeDuplicateRequest    QueryResponseCode = "94"
	QueryCodeInvalidChecksum     QueryResponseCode = "97"
	QueryCodeOtherErrors         QueryResponseCode = "99"
)

// ToString returns the response code as a two-digit string
func (code QueryResponseCode) ToString() string {
	return string(code)
}

// Message returns a detailed message for the response code
func (code QueryResponseCode) Message() string {
	switch code {
	case QueryCodeRequestSuccess:
		return "Request successful"
	case QueryCodeInvalidConnectionID:
		return "Invalid connection identifier"
	case QueryCodeInvalidDataFormat:
		return "Invalid data format sent"
	case QueryCodeTransactionNotFound:
		return "Transaction not found"
	case QueryCodeDuplicateRequest:
		return "Duplicate request within the API's limited time frame"
	case QueryCodeInvalidChecksum:
		return "Invalid checksum"
	case QueryCodeOtherErrors:
		return "Other errors"
	default:
		return fmt.Sprintf("Unknown response code: %s", code)
	}
}

func (code QueryResponseCode) IsRequestSuccess() bool {
	return code == QueryCodeRequestSuccess
}

func (code QueryResponseCode) IsInvalidConnectionID() bool {
	return code == QueryCodeInvalidConnectionID
}

func (code QueryResponseCode) IsInvalidDataFormat() bool {
	return code == QueryCodeInvalidDataFormat
}

func (code QueryResponseCode) IsTransactionNotFound() bool {
	return code == QueryCodeTransactionNotFound
}

func (code QueryResponseCode) IsDuplicateRequest() bool {
	return code == QueryCodeDuplicateRequest
}

func (code QueryResponseCode) IsInvalidChecksum() bool {
	return code == QueryCodeInvalidChecksum
}

func (code QueryResponseCode) IsOtherErrors() bool {
	return code == QueryCodeOtherErrors
}
