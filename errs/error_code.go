package errs

var (
	Success = New(0000, "success")

	// 通用错误
	ParamError          = New(0001, "param error")
	DbError             = New(0002, "db error")
	AmountCurrencyError = New(0003, "amount currency error")
	SystemError         = New(0004, "server error")
	RpcRemoteError      = New(0005, "call remote rpc error")
	HttpRemoteError     = New(0006, "call http error")
	TxnNotExist         = New(0007, "txn not exist")
)

// internal system error
var (
	DbDuplicatedError   = New(90000000, "db creation duplicated")
	JsonError           = New(90000001, "json operation error")
	KeyNotFoundError    = New(90000002, "key not found error")
	InvalidStateMachine = New(90000003, "invalid status machine")
	NotifyFail          = New(90000004, "notify fail")
)
