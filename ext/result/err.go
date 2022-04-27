//@File     err.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

import "fmt"

/**
常用通用固定错误
*/

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
