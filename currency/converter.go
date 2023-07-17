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

// CurrencyFraction 指定currency的rule，在第三方转义场景使用
func CurrencyFraction(rules map[string]int32, currency string) (int32, errs.Error) {
	fraction, ok := rules[currency]
	if !ok {
		return 0, errs.AmountCurrencyError.ResetMsg("unknown currency")
	}
	return fraction, nil
}

func ConvertMicroToMinor(micro int64, currency string) (string, errs.Error) {
	fraction, err := konoCurrencyFraction(currency)
	if err != nil {
		return "", err
	}
	return decimal.New(micro, -6).StringFixedBank(fraction), nil
}

func ConvertMicroToMinorDecimal(micro int64, currency string) (decimal.Decimal, errs.Error) {
	fraction, err := konoCurrencyFraction(currency)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal.New(micro, -6).RoundBank(fraction), nil
}
