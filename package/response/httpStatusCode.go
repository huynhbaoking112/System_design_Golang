package response

const (
	ErrCodeSusscess     = 2001 //success
	ErrCodeParamInvalid = 2003 //Email is invalid
)

// message

var msg = map[int]string{
	ErrCodeSusscess:     "success",
	ErrCodeParamInvalid: "Email is invalid",
}
