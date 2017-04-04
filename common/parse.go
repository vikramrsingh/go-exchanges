package common

import (
	"strconv"

	"github.com/crypto-bank/proto/currency"
)

// ParseIFloat64 - Parse interface to string to float.
func ParseIFloat64(i interface{}) (_ float64, err error) {
	return strconv.ParseFloat(i.(string), 64)
}

// ParseIInt64 - Parse interface to string to float.
func ParseIInt64(i interface{}) (_ int64, err error) {
	return strconv.ParseInt(i.(string), 10, 64)
}

// ParseIVolume - Parse interface to string to float and to currency volume.
func ParseIVolume(cur *currency.Currency, i interface{}) (_ *currency.Volume, err error) {
	amount, err := ParseIFloat64(i)
	if err != nil {
		return
	}
	return &currency.Volume{
		Amount:   amount,
		Currency: cur,
	}, nil
}
