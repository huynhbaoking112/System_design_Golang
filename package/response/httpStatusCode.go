package response

const (
	ErrCodeSusscess     = 2001 //success
	ErrCodeParamInvalid = 2003 //Email is invalid

	ErrInvalidToken = 30001 //Token is invalid

	// Register Code
	ErrCodeUserHasExists = 50001 // user has already registered
)

// message

var msg = map[int]string{
	ErrCodeSusscess:      "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "User has already registered",
}
