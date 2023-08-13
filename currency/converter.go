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

func ConvertMinorFloat2Micro(m float64) int64 {
	value := decimal.NewFromFloat(m)
	return value.Mul(decimal.New(1, 6)).IntPart()
}

// ConvertMicroToExternalMinor 指定external currency的rule，在第三方转义场景使用
func ConvertMicroToExternalMinor(micro int64, currency string, fractionRules map[string]int32) (string, errs.Error) {
	fraction, ok := fractionRules[currency]
	if !ok {
		return "", errs.AmountCurrencyError.ResetMsg("unknown currency")
	}
	return decimal.New(micro, -6).StringFixedBank(fraction), nil
}

func ConvertMicroToMinor(micro int64, currency string) (string, errs.Error) {
	fraction, err := konoCurrencyFraction(currency)
	if err != nil {
		return "", err
	}
	return decimal.New(micro, -6).StringFixedBank(fraction), nil
}

func ConvertMicroToMinorFloat(micro int64) float64 {
	return decimal.New(micro, -6).InexactFloat64()
}

func ConvertMicroToMinorDecimal(micro int64, currency string) (decimal.Decimal, errs.Error) {
	fraction, err := konoCurrencyFraction(currency)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal.New(micro, -6).RoundBank(fraction), nil
}
