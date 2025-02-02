package response

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Remark     string `json:"remark"`
}

func (e *Error) Error() string {
	return e.Remark
}

func New(statusCode int, code, remark string) *Error {
	return &Error{
		Code:   code,
		Remark: remark,
	}
}

var (
	ErrDuplicateIdentityNumber = New(http.StatusBadRequest, "DUPLICATE_IDENTITY", "identity number already exists")
	ErrDuplicatePhoneNumber    = New(http.StatusBadRequest, "DUPLICATE_PHONE", "phone number already exists")
	ErrDuplicateAccountNumber  = New(http.StatusBadRequest, "DUPLICATE_ACCOUNT", "account number already exists")
)
