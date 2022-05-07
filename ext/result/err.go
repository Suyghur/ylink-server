//@File     err.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

import "fmt"

/**
常用通用固定错误
*/

const (
	Ok                 int64 = 0
	ServerCommonError  int64 = 10001
	RequestParamError  int64 = 10002
	TokenExpireError   int64 = 10003
	TokenParseError    int64 = 10004
	TokenGenerateError int64 = 10005
	DbError            int64 = 10006
)

var message map[int64]string

func init() {
	message = make(map[int64]string)
	message[Ok] = "success"
	message[ServerCommonError] = "some unknown error has occurred"
	message[RequestParamError] = "request params error"
	message[TokenExpireError] = "the token is invalid, please re-authenticate"
	message[TokenParseError] = "failed to parse token"
	message[TokenGenerateError] = "failed to generate token"
	message[DbError] = "database is busy, please try again late"
}

func MapErrMsg(code int64) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "some unknown error has occurred"
	}
}

func IsCodeErr(code int64) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}

type CodeError struct {
	errCode int64
	errMsg  string
}

func (e *CodeError) GetErrCode() int64 {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("code:%d，msg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(code int64, msg string) *CodeError {
	return &CodeError{errCode: code, errMsg: msg}
}
func NewErrCode(errCode int64) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: ServerCommonError, errMsg: errMsg}
}
