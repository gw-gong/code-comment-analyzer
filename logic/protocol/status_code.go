package protocol

const (
	Success int = iota
)

const (
	ErrorCodeAuthenticating int = 1000 + iota
	ErrorCodeAuthorizing
	ErrorCodeMissingUserId
	ErrorCodeRPCCallFail
	ErrorCodeMustBeGet
	ErrorCodeMustBePost
)
