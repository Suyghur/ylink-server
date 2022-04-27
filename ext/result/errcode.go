//@File     errcode.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

const (
	Ok                 int64 = 0
	ServerCommonError  int64 = 10001
	RequestParamError  int64 = 10002
	TokenExpireError   int64 = 10003
	TokenParseError    int64 = 10004
	TokenGenerateError int64 = 10005
	DbError            int64 = 10006
)

/**
前2位代表业务,后3位代表具体功能
**/
const ()
