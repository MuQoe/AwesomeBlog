package errcode

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type Error struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Detail []string `json:"detail"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		errmsg := fmt.Sprintf("错误码 %d 已经存在, 请更换一个！", code)
		zap.S().Errorf("错误码 %d 已经存在, 请更换一个！", code)
		panic(errmsg)
	}
	codes[code] = msg
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Msg)
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Detail = []string{}
	for _, d := range details {
		newError.Detail = append(newError.Detail, d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeOut.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
