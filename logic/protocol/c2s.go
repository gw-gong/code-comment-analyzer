package protocol

import "strconv"

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

// *********** project tree ***********

type FileNode struct {
	Label    string     `json:"label"`
	Value    string     `json:"value,omitempty"`
	Children []FileNode `json:"children,omitempty"`
}

// *********** get user info ***********

type TableInfo struct {
	NickName   string `json:"nickname"`
	Email      string `json:"email"`
	DateJoined string `json:"date_joined"`
}

type GetUserInfoResponse struct {
	ProfilePicture string      `json:"profile_picture"`
	TableInfo      []TableInfo `json:"tableInfo"`
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

// *********** get file/project upload record ***********

const GetKeyOperatingRecordId = "operating_record_id"
const GetUserOperatingRecords = "operating_record_id"

func OpIDTransformStr2Int64(opID string) int64 {
	result, err := strconv.ParseInt(opID, 10, 64)
	if err != nil {
		return 0
	}
	return result
}

//======change_password----------------
type ChangePasswordRequest struct {
	OldPassword      string `json:"old_password"`
	NewPassword      string `json:"new_password"`
	AgainNewPassword string `json:"again_new_password"`
}

//=====delete_operating_record------------
// Extract operation ID from the request body
type DeleteOperatingRecordRequest struct {
	ID int64 `json:"id"`
}

//=====update_user_info------------
type UpdateInfo struct {
	Nickname         string `json:"nickname"`
	AgainNewPassword string `json:"again_new_password"`
}
type UpdateUserInfoRequest struct {
	Nickname         string `json:"nickname"`
	AgainNewPassword string `json:"again_new_password"`
}
