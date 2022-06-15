//@File     err.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

import "fmt"

/**
常用通用固定错误
*/

const (
	Ok                 int32 = 0
	ServerCommonError  int32 = 10001
	RequestParamError  int32 = 10002
	TokenExpireError   int32 = 10003
	TokenParseError    int32 = 10004
	TokenGenerateError int32 = 10005
	DbError            int32 = 10006
)

var message map[int32]string

func init() {
	message = make(map[int32]string)
	message[Ok] = "success"
	message[ServerCommonError] = "some unknown error has occurred"
	message[RequestParamError] = "request params error"
	message[TokenExpireError] = "the token is invalid, please re-authenticate"
	message[TokenParseError] = "failed to parse token"
	message[TokenGenerateError] = "failed to generate token"
	message[DbError] = "database is busy, please try again late"
}

func MapErrMsg(code int32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "some unknown error has occurred"
	}
}

func IsCodeErr(code int32) bool {
	if _, ok := message[code]; ok {
		return true
	} else {
		return false
	}
}

type CodeError struct {
	errCode int32
	errMsg  string
}

func (e *CodeError) GetErrCode() int32 {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("code:%d，msg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(code int32, msg string) *CodeError {
	return &CodeError{errCode: code, errMsg: msg}
}
func NewErrCode(errCode int32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: ServerCommonError, errMsg: errMsg}
}
