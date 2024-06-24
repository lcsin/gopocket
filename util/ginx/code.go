package ginx

//go:generate stringer -type=ErrCode -linecomment
type ErrCode int

const (
	ErrSuccess        ErrCode = 0    // success
	ErrInternalServer ErrCode = -1   // internal server error
	ErrBadRequest     ErrCode = -400 // invalid argument
	ErrUnauthorized   ErrCode = -401 // unauthorized
	ErrNotFound       ErrCode = -404 // not found
)
