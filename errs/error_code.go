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
)
