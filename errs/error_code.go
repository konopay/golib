package errs

var (
	Success = New(0, "success")

	// 通用错误
	ParamError          = New(1, "param error")
	DbError             = New(2, "db error")
	AmountCurrencyError = New(3, "amount currency error")
	SystemError         = New(4, "server error")
	RpcRemoteError      = New(5, "call remote rpc error")
	HttpRemoteError     = New(6, "call http error")
	TxnNotExist         = New(7, "txn not exist")
	InsufficientBalance = New(8, "insufficient balance")
	ChannelApiError     = New(9, "channel api error")
	ChannelBizError     = New(10, "channel biz error")
	InvalidUserPI       = New(11, "invalid user instrument details")
	EncryptError        = New(12, "Encrypt Error")
	DecryptError        = New(13, "Decrypt Error")
	TxnExpired          = New(14, "transaction expired")
)

// internal system error
var (
	DbDuplicatedError    = New(90000000, "db creation duplicated")
	JsonError            = New(90000001, "json operation error")
	KeyNotFoundError     = New(90000002, "key not found error")
	InvalidStateMachine  = New(90000003, "invalid status machine")
	NotifyFail           = New(90000004, "notify fail")
	NoValidChannelFound  = New(90000005, "channel not found")
	DataNotFoundError    = New(90000006, "data not found in db")
	ShouldNotHappenError = New(90000007, "internal error should not happen")
	UnknownChannelStatus = New(90000008, "unknown channel status")
	DuplicatedRequest    = New(90000009, "duplicated request")
)
