package helper

import "strconv"

func ParseAmount(amount string) int64 {
	parsedAmount, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return 0
	}
	return parsedAmount
}

func ParseInt64(value string) int64 {
	parsedValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return parsedValue
}

func ParseInt32(value string) int32 {
	parsedValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0
	}
	return int32(parsedValue)
}
