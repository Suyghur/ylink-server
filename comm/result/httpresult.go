//@File     httpresult.go
//@Time     2022/04/26
//@Author   #Suyghur,

package result

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		result := Success(resp)
		httpx.WriteJson(w, http.StatusOK, result)
	} else {
		code := ServerCommonError
		msg := "some unknown error has occurred"

		cause := errors.Cause(err)
		if e, ok := cause.(*CodeError); ok {
			code = e.GetErrCode()
			msg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(cause); ok {
				grpcCode := int32(gStatus.Code())
				if IsCodeErr(grpcCode) {
					code = grpcCode
					msg = gStatus.Message()
				} else {

				}
			}
		}

		logx.WithContext(r.Context()).Errorf("[API ERROR] : %+v", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(code, msg))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	logx.WithContext(r.Context()).Errorf("[API ERROR] : %+v", err)
	msg := fmt.Sprintf("%s ,%s", MapErrMsg(RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(RequestParamError, msg))
}
