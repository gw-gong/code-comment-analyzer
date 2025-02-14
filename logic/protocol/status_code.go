package protocol

const (
	StatusSuccess int = iota
)

const (
	ErrorCodeAuthenticating int = 1000 + iota
	ErrorCodeAuthorizing
	ErrorCodeMissingUserId
	ErrorCodeAnalyzeFileFailed

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

	ErrorCodeInvalidRequest
	ErrorCodeInvalidPassword

	ErrorCodeInvalidID
	ErrorCodeBadRequest
	ErrorCodeUpdateUserInfoFailed
	ErrorCodeUpdateUserAvatarFailed

	ErrorCodeGetFileUploadRecordFailed
	ErrorCodeGetFileContentFailed
	ErrorCodeDeleteOperatingRecordFailed
	ErrorCodeGetUserOperatingRecordsFailed
	ErrorCodeGetUserInfoFailed
	ErrorCodeGetUserProfilePictureFailed
	ErrorCodeGetProjectUploadRecordFailed
	ErrorCodeChangePasswordFailed
	ErrorCodeInternalError
)
