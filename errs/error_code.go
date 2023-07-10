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
)

// internal system error
var (
	DbDuplicatedError   = New(90000000, "db creation duplicated")
	JsonError           = New(90000001, "json operation error")
	KeyNotFoundError    = New(90000002, "key not found error")
	InvalidStateMachine = New(90000003, "invalid status machine")
	NotifyFail          = New(90000004, "notify fail")
	NoValidChannelFound = New(90000005, "channel not found")
)
