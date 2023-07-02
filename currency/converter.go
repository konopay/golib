package currency

import (
	"github.com/konopay/golib/errs"
	"github.com/shopspring/decimal"
)

func ConvertMinor2Micro(m string) (int64, errs.Error) {
	value, err := decimal.NewFromString(m)
	if err != nil {
		return 0, errs.AmountCurrencyError
	}
	return value.Mul(decimal.New(1, 6)).IntPart(), nil
}

func ConvertMicroToMinor(micro int64, friction int32) string {
	return decimal.New(micro, -6).StringFixedBank(friction)
}

func ConvertMicroToMinorDecimal(micro int64, friction int32) decimal.Decimal {
	return decimal.New(micro, -6).RoundBank(friction)
}

func CurrencyFriction(rules map[string]int32, currency string) (int32, errs.Error) {
	friction, ok := rules[currency]
	if !ok {
		return 0, errs.AmountCurrencyError.ResetMsg("unknown currency")
	}
	return friction, nil
}

// TODO: add standard rule mapping
func ISOCurrencyFriction(currency string) (int32, errs.Error) {
	return 0, errs.AmountCurrencyError
}
