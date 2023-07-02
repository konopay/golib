package errs

var (
	Success = New(0000, "success")

	// 通用错误
	ParamError      = New(0001, "param error")
	DbError         = New(0002, "db error")
	SystemError     = New(0003, "server error")
	RpcRemoteError  = New(0004, "call remote rpc error")
	HttpRemoteError = New(0005, "call http error")
)
