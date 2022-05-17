//@File     responsebean.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

type ResponseBean struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *ResponseBean {
	if data == nil {
		data = map[string]interface{}{}
	}
	return &ResponseBean{0, "success", data}
}

func Error(code int64, msg string) *ResponseBean {
	return &ResponseBean{code, msg, map[string]interface{}{}}
}
