package protocol

// *********** sign in ***********

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UID      uint64 `json:"uid"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

// *********** sign up ***********

// *********** sign out ***********

// *********** analyze file ***********

type AnalyzeFileRequest struct {
	Language    string `json:"language"`
	FileContent string `json:"fileContent"`
}

type AnalyzeFileResponse map[string]interface{}

// *********** file to string ***********

type File2StringResponse struct {
	Language    string `json:"language"`
	FileContent string `json:"fileContent"`
}

type SignupRequest struct {
	Email         string `json:"email"`          // 用户邮箱
	Password      string `json:"password"`       // 用户密码
	PasswordAgain string `json:"password_again"` // 确认密码
}

type SignupResponse struct {
	UID      uint64 `json:"uid"`      // 用户ID
	Email    string `json:"email"`    // 用户邮箱
	Nickname string `json:"nickname"` // 用户昵称
}
