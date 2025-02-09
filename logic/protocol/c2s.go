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

type SignupRequest struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordAgain string `json:"password_again"`
}

type SignupResponse struct {
	UID      uint64 `json:"uid"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

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

// *********** upload and get tree ***********

type FileNode struct {
	Label    string     `json:"label"`
	Value    string     `json:"value,omitempty"`
	Children []FileNode `json:"children,omitempty"`
}

// *********** get user info ***********

type GetUserInfoResponse struct {
	NickName   string `json:"nick_name"`
	Email      string `json:"email"`
	DateJoined string `json:"date_joined"`
}

// *********** get user profile picture ***********

type GetUserProfilePictureResponse struct {
	ProfilePicture *string `json:"profile_picture"` // 为了和重构前返回的结果一致
	Text           string  `json:"text"`
}

// *********** read file ***********

type ReadFileRequest struct {
	Path string `json:"path"`
}

type ReadFileResponse struct {
	FileContent string `json:"fileContent"`
}
