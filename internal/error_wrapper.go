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

	ErrorInvalidRequest                  = regErrorForFE(400000, "invalid request")
	ErrorInvalidRequestUserName          = regErrorForFE(400001, "invalid request, username is required, or character more then 30")
	ErrorInvalidRequestPassword          = regErrorForFE(400002, "invalid request, password is required, or character more then 30")
	ErrorInvalidRequestCompanyName       = regErrorForFE(400002, "invalid request, company name is required, or character less 3 and more then 30")
	ErrorInvalidRequestPhoneNumber       = regErrorForFE(400002, "invalid request, phone number is required, or character less 8 and more then 16")
	ErrorInvalidRequestAddress           = regErrorForFE(400002, "invalid request, address is required, or character less 10 and more then 50")
	ErrorInvalidRequestJobTitle          = regErrorForFE(400002, "invalid request, job title is required, or job title not either manager, director, staff")
	ErrorInvalidRequestEmail             = regErrorForFE(400002, "invalid request, email is required, or wrong format email, or character less 5 or more then 255")
	ErrorInvalidLogin                    = regErrorForFE(400003, "login invalid")
	ErrMissingAuthorizationHeader        = regErrorForFE(400004, "invalid authorization")
	ErrInvalidTokenFormat                = regErrorForFE(400005, "invalid token format")
	ErrInvalidToken                      = regErrorForFE(400006, "invalid token")
	ErrNoModifyUpdate                    = regErrorForFE(400007, "error no data modify update")
	ErrDataCompanyNotFound               = regErrorForFE(400008, "error company not found")
	ErrorInvalidParameterID              = regErrorForFE(400009, "invalid parameter id")
	ErrorInvalidDataNotFound             = regErrorForFE(400009, "error, invalid data not found")
	ErrorInvalidInsertedPhoneDuplicated  = regErrorForFE(400010, "error, data phone number duplicated")
	ErrorInvalidDuplicateCompanyAndEmail = regErrorForFE(400011, "error, email already registered to other company-duplicated")
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
