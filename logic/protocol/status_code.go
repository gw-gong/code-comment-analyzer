package protocol

const (
	StatusSuccess int = iota
)

const (
	ErrorCodeAuthenticating int = 1000 + iota
	ErrorCodeAuthorizing
	ErrorCodeMissingUserId
	ErrorCodeRPCCallFail

	ErrorCodeMustBeGet
	ErrorCodeMustBePost

	ErrorCodeParseRequestFailed
	ErrorCodeParamError
	ErrorCodeRegisteredEmail
	ErrorCodeInternalServerError

	ErrorCodeLanguageNotSupported

	ErrorCodeFileTooLarge
	ErrorCodeFileNotFound

	ErrorCodeCreatePathFailed
	ErrorCodeSaveFileFailed
	ErrorCodeUnzipFailed
)
