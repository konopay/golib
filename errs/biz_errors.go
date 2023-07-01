package errs

var (
	Success = New(0000, "success")

	// 通用错误
	DbError        = New(0001, "db error")
	SystemError    = New(0002, "server error")
	RpcRemoteError = New(0003, "call remote error")
	ParamError     = New(0004, "param error")
)
