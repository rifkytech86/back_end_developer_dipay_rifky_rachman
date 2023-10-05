package internal

type ErrorMobile string

var (
	errorCode   = make(map[int64]string)
	errorString = make(map[string]int64)
)

var (
	ErrorInternalServer        = regErrorForFE(500000, "internal server error")
	ErrorInternalReadFile      = regErrorForFE(500001, "internal error read file")
	ErrorInternalGenerateToken = regErrorForFE(500002, "login invalid")

	ErrorInvalidRequest           = regErrorForFE(400000, "invalid request")
	ErrorInvalidRequestUserName   = regErrorForFE(400001, "invalid request, username is required")
	ErrorInvalidRequestPassword   = regErrorForFE(400002, "invalid request, password is required")
	ErrorInvalidLogin             = regErrorForFE(400003, "login invalid")
	ErrMissingAuthorizationHeader = regErrorForFE(400004, "invalid authorization")
	ErrInvalidTokenFormat         = regErrorForFE(400005, "invalid token format")
	ErrInvalidToken               = regErrorForFE(400006, "invalid token")
	ErrNoModifyUpdate             = regErrorForFE(400007, "error no data modify update")
	ErrDataCompanyNotFound        = regErrorForFE(400008, "error company not found")
	ErrorInvalidParameterID       = regErrorForFE(400009, "invalid parameter id")
	ErrorInvalidDataNotFound      = regErrorForFE(400009, "error, invalid data not found")
)

func regErrorForFE(code int64, msg string) ErrorMobile {
	errorString[msg] = code
	errorCode[code] = msg
	return ErrorMobile(msg)
}

func (e ErrorMobile) String() string {
	return string(e)
}

func (e ErrorMobile) GetCode() int64 {
	code, exists := errorString[string(e)]
	if !exists {
		return 0
	}
	return code
}

func GetCodeByString(messageError string) int64 {
	code, exists := errorString[messageError]
	if !exists {
		return 0
	}
	return code
}

func GetMsg(code int64) string {
	msg, exists := errorCode[code]
	if !exists {
		return ""
	}
	return msg
}
