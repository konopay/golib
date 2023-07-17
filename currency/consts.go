package currency

import "github.com/konopay/golib/errs"

const (
	AED = "AED"
	AFN = "AFN"
	ALL = "ALL"
	AMD = "AMD"
	ANG = "ANG"
	AOA = "AOA"
	ARS = "ARS"
	AUD = "AUD"
	AWG = "AWG"
	AZN = "AZN"
	BAM = "BAM"
	BBD = "BBD"
	BDT = "BDT"
	BGN = "BGN"
	BHD = "BHD"
	BIF = "BIF"
	BMD = "BMD"
	BND = "BND"
	BOB = "BOB"
	BRL = "BRL"
	BSD = "BSD"
	BTN = "BTN"
	BWP = "BWP"
	BYN = "BYN"
	BYR = "BYR"
	BZD = "BZD"
	CAD = "CAD"
	CDF = "CDF"
	CHF = "CHF"
	CLF = "CLF"
	CLP = "CLP"
	CNY = "CNY"
	COP = "COP"
	CRC = "CRC"
	CUC = "CUC"
	CUP = "CUP"
	CVE = "CVE"
	CZK = "CZK"
	DJF = "DJF"
	DKK = "DKK"
	DOP = "DOP"
	DZD = "DZD"
	EEK = "EEK"
	EGP = "EGP"
	ERN = "ERN"
	ETB = "ETB"
	EUR = "EUR"
	FJD = "FJD"
	FKP = "FKP"
	GBP = "GBP"
	GEL = "GEL"
	GGP = "GGP"
	GHC = "GHC"
	GHS = "GHS"
	GIP = "GIP"
	GMD = "GMD"
	GNF = "GNF"
	GTQ = "GTQ"
	GYD = "GYD"
	HKD = "HKD"
	HNL = "HNL"
	HRK = "HRK"
	HTG = "HTG"
	HUF = "HUF"
	IDR = "IDR"
	ILS = "ILS"
	IMP = "IMP"
	INR = "INR"
	IQD = "IQD"
	IRR = "IRR"
	ISK = "ISK"
	JEP = "JEP"
	JMD = "JMD"
	JOD = "JOD"
	JPY = "JPY"
	KES = "KES"
	KGS = "KGS"
	KHR = "KHR"
	KMF = "KMF"
	KPW = "KPW"
	KRW = "KRW"
	KWD = "KWD"
	KYD = "KYD"
	KZT = "KZT"
	LAK = "LAK"
	LBP = "LBP"
	LKR = "LKR"
	LRD = "LRD"
	LSL = "LSL"
	LTL = "LTL"
	LVL = "LVL"
	LYD = "LYD"
	MAD = "MAD"
	MDL = "MDL"
	MGA = "MGA"
	MKD = "MKD"
	MMK = "MMK"
	MNT = "MNT"
	MOP = "MOP"
	MUR = "MUR"
	MVR = "MVR"
	MWK = "MWK"
	MXN = "MXN"
	MYR = "MYR"
	MZN = "MZN"
	NAD = "NAD"
	NGN = "NGN"
	NIO = "NIO"
	NOK = "NOK"
	NPR = "NPR"
	NZD = "NZD"
	OMR = "OMR"
	PAB = "PAB"
	PEN = "PEN"
	PGK = "PGK"
	PHP = "PHP"
	PKR = "PKR"
	PLN = "PLN"
	PYG = "PYG"
	QAR = "QAR"
	RON = "RON"
	RSD = "RSD"
	RUB = "RUB"
	RUR = "RUR"
	RWF = "RWF"
	SAR = "SAR"
	SBD = "SBD"
	SCR = "SCR"
	SDG = "SDG"
	SEK = "SEK"
	SGD = "SGD"
	SHP = "SHP"
	SKK = "SKK"
	SLL = "SLL"
	SOS = "SOS"
	SRD = "SRD"
	SSP = "SSP"
	STD = "STD"
	SVC = "SVC"
	SYP = "SYP"
	SZL = "SZL"
	THB = "THB"
	TJS = "TJS"
	TMT = "TMT"
	TND = "TND"
	TOP = "TOP"
	TRL = "TRL"
	TRY = "TRY"
	TTD = "TTD"
	TWD = "TWD"
	TZS = "TZS"
	UAH = "UAH"
	UGX = "UGX"
	USD = "USD"
	UYU = "UYU"
	UZS = "UZS"
	VEF = "VEF"
	VND = "VND"
	VUV = "VUV"
	WST = "WST"
	XAF = "XAF"
	XAG = "XAG"
	XAU = "XAU"
	XCD = "XCD"
	XDR = "XDR"
	XOF = "XOF"
	XPF = "XPF"
	YER = "YER"
	ZAR = "ZAR"
	ZMW = "ZMW"
	ZWD = "ZWD"
	ZWL = "ZWL"
	MRO = "MRO"
	CNH = "CNH"
	VES = "VES"
)

type Currency struct {
	Decimal  string
	Thousand string
	Fraction int32
	Grapheme string
	Template string
}

var konoCurrencyMap = map[string]Currency{
	TWD: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "NT$", Template: "$1"},
	IDR: {Decimal: ",", Thousand: ".", Fraction: 0, Grapheme: "Rp", Template: "$1"},
	AED: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u062f.\u0625", Template: "1 $"},
	AFN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u060b", Template: "1 $"},
	ALL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "L", Template: "$1"},
	AMD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u0564\u0580.", Template: "1 $"},
	ANG: {Decimal: ",", Thousand: ".", Fraction: 2, Grapheme: "\u0192", Template: "$1"},
	AOA: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Kz", Template: "1$"},
	ARS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	AUD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	AWG: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u0192", Template: "1$"},
	AZN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20bc", Template: "$1"},
	BAM: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "KM", Template: "$1"},
	BBD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	BDT: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u09f3", Template: "$1"},
	BGN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u043b\u0432", Template: "$1"},
	BHD: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\u062f.\u0628", Template: "1 $"},
	BIF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "Fr", Template: "1$"},
	BMD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	BND: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	BOB: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Bs.", Template: "$1"},
	BRL: {Decimal: ",", Thousand: ".", Fraction: 2, Grapheme: "R$", Template: "$1"},
	BSD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	BTN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Nu.", Template: "1$"},
	BWP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "P", Template: "$1"},
	BYN: {Decimal: ",", Thousand: " ", Fraction: 2, Grapheme: "p.", Template: "1 $"},
	BYR: {Decimal: ",", Thousand: " ", Fraction: 0, Grapheme: "p.", Template: "1 $"},
	BZD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "BZ$", Template: "$1"},
	CAD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	CDF: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "FC", Template: "1$"},
	CHF: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "CHF", Template: "1 $"},
	CLF: {Decimal: ",", Thousand: ".", Fraction: 4, Grapheme: "UF", Template: "$1"},
	CLP: {Decimal: ",", Thousand: ".", Fraction: 0, Grapheme: "$", Template: "$1"},
	CNY: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u5143", Template: "1 $"},
	CNH: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u5143", Template: "1 $"},
	COP: {Decimal: ",", Thousand: ".", Fraction: 2, Grapheme: "$", Template: "$1"},
	CRC: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a1", Template: "$1"},
	CUC: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "1$"},
	CUP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$MN", Template: "$1"},
	CVE: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "1$"},
	CZK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "K\u010d", Template: "1 $"},
	DJF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "Fdj", Template: "1 $"},
	DKK: {Decimal: ",", Thousand: ".", Fraction: 2, Grapheme: "kr", Template: "$ 1"},
	DOP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "RD$", Template: "$1"},
	DZD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u062f.\u062c", Template: "1 $"},
	EEK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "kr", Template: "$1"},
	EGP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	ERN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Nfk", Template: "1 $"},
	ETB: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Br", Template: "1 $"},
	EUR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20ac", Template: "$1"},
	FJD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	FKP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	GBP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	GEL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u10da", Template: "1 $"},
	GGP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	GHC: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a2", Template: "$1"},
	GHS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20b5", Template: "$1"},
	GIP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	GMD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "D", Template: "1 $"},
	GNF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "FG", Template: "1 $"},
	GTQ: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Q", Template: "$1"},
	GYD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	HKD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	HNL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "L", Template: "$1"},
	HRK: {Decimal: ",", Thousand: ".", Fraction: 2, Grapheme: "kn", Template: "1 $"},
	HTG: {Decimal: ",", Thousand: ".", Fraction: 2, Grapheme: "G", Template: "1 $"},
	HUF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "Ft", Template: "$1"},

	ILS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20aa", Template: "$1"},
	IMP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	INR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20b9", Template: "$1"},
	IQD: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\u062f.\u0639", Template: "1 $"},
	IRR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\ufdfc", Template: "1 $"},
	ISK: {Decimal: ",", Thousand: ".", Fraction: 0, Grapheme: "kr", Template: "$1"},
	JEP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	JMD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "J$", Template: "$1"},
	JOD: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\u062f.\u0625", Template: "1 $"},
	JPY: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "\u00a5", Template: "$1"},
	KES: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "KSh", Template: "$1"},
	KGS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u0441\u043e\u043c", Template: "$1"},
	KHR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u17db", Template: "$1"},
	KMF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "CF", Template: "$1"},
	KPW: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "\u20a9", Template: "$1"},
	KRW: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "\u20a9", Template: "$1"},
	KWD: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\u062f.\u0643", Template: "1 $"},
	KYD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	KZT: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20b8", Template: "$1"},
	LAK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20ad", Template: "$1"},
	LBP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	LKR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a8", Template: "$1"},
	LRD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	LSL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "L", Template: "$1"},
	LTL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Lt", Template: "$1"},
	LVL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Ls", Template: "1 $"},
	LYD: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\u062f.\u0644", Template: "1 $"},
	MAD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u062f.\u0645", Template: "1 $"},
	MDL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "lei", Template: "1 $"},
	MKD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u0434\u0435\u043d", Template: "$1"},
	MMK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "K", Template: "$1"},
	MNT: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20ae", Template: "$1"},
	MOP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "P", Template: "1 $"},
	MUR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a8", Template: "$1"},
	MVR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "MVR", Template: "1 $"},
	MWK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "MK", Template: "$1"},
	MXN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	MYR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "RM", Template: "$1"},
	MZN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "MT", Template: "$1"},
	NAD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	NGN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a6", Template: "$1"},
	NIO: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "C$", Template: "$1"},
	NOK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "kr", Template: "1 $"},
	NPR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a8", Template: "$1"},
	NZD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	OMR: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\ufdfc", Template: "1 $"},
	PAB: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "B/.", Template: "$1"},
	PEN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "S/", Template: "$1"},
	PGK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "K", Template: "1 $"},
	PHP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20b1", Template: "$1"},
	PKR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a8", Template: "$1"},
	PLN: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "z\u0142", Template: "1 $"},
	PYG: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "Gs", Template: "1$"},
	QAR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\ufdfc", Template: "1 $"},
	RON: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "lei", Template: "$1"},
	RSD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u0414\u0438\u043d.", Template: "$1"},
	RUB: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20bd", Template: "1 $"},
	RUR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20bd", Template: "1 $"},
	RWF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "FRw", Template: "1 $"},
	SAR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\ufdfc", Template: "1 $"},
	SBD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	SCR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a8", Template: "$1"},
	SDG: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	SEK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "kr", Template: "1 $"},
	SGD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	SHP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	SKK: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Sk", Template: "$1"},
	SLL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Le", Template: "1 $"},
	SOS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Sh", Template: "1 $"},
	SRD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	SSP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "1 $"},
	STD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Db", Template: "1 $"},
	SVC: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a1", Template: "$1"},
	SYP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "1 $"},
	SZL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u00a3", Template: "$1"},
	THB: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u0e3f", Template: "$1"},
	TJS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "SM", Template: "1 $"},
	TMT: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "T", Template: "1 $"},
	TND: {Decimal: ".", Thousand: ",", Fraction: 3, Grapheme: "\u062f.\u062a", Template: "1 $"},
	TOP: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "T$", Template: "$1"},
	TRL: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20a4", Template: "$1"},
	TRY: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20ba", Template: "$1"},
	TTD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "TT$", Template: "$1"},

	TZS: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "TSh", Template: "$1"},
	UAH: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\u20b4", Template: "1 $"},
	UGX: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "USh", Template: "1 $"},
	USD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	UYU: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$U", Template: "$1"},
	UZS: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "so\u2019m", Template: "$1"},
	VEF: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Bs", Template: "$1"},
	VES: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Bs.", Template: "$ 1"},
	VND: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "\u20ab", Template: "1 $"},
	VUV: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "Vt", Template: "$1"},
	WST: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "T", Template: "1 $"},
	XAF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "Fr", Template: "1 $"},
	XAG: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "oz t", Template: "1 $"},
	XAU: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "oz t", Template: "1 $"},
	XCD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "$", Template: "$1"},
	XDR: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "SDR", Template: "1 $"},
	YER: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "\ufdfc", Template: "1 $"},
	ZAR: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "R", Template: "$1"},
	ZMW: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "ZK", Template: "$1"},
	ZWD: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "Z$", Template: "$1"},

	MRO: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "MRO", Template: "$1"},
	XOF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "XOF", Template: "$1"},
	XPF: {Decimal: ".", Thousand: ",", Fraction: 0, Grapheme: "XPF", Template: "$1"},
	MGA: {Decimal: ".", Thousand: ",", Fraction: 2, Grapheme: "MGA", Template: "$1"},
}

func konoCurrencyFraction(currency string) (int32, errs.Error) {
	c, ok := konoCurrencyMap[currency]
	if !ok {
		return 0, errs.AmountCurrencyError.ResetMsg("currency not found")
	}

	return c.Fraction, nil
}
