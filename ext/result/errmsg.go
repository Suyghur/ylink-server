//@File     errmsg.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

var message map[int64]string

func init() {
	message = make(map[int64]string)
	message[Ok] = "success"
	message[ServerCommonError] = "some unknown error has occurred"
	message[RequestParamError] = "request params error"
	message[TokenExpireError] = "the token is invalid, please re-authenticate"
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
