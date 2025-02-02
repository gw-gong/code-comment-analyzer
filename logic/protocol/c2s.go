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
