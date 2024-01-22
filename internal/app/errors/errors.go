package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("没找到相关记录")
	ErrServerErr      = errors.New("服务内部错误")

	ErrApiRequestFail = errors.New("接口请求失败")
)

type Error struct {
	error
}

func E(err error) *Error {
	return &Error{
		err,
	}
}
func (err *Error) Stack() string {
	return fmt.Sprintf("%+v", err.error)
}
func (err *Error) Message() string {
	return err.Error()
}
func (err *Error) MessageCode() (string, int) {
	return err.Message(), err.Code()
}
func (err *Error) Code() int {
	switch true {
	case Is(err.error, ErrRecordNotFound):
		return 10004
	case Is(err.error, ErrServerErr):
		return 10005
	//case ErrSqlFindErr:
	//	return 10006
	//case ErrRequestDataErr:
	//	return 10007
	//case ErrPhoneNum86:
	//	return 10008
	//case ErrNoBindingWeixin:
	//	return 10009
	//case ErrNoBindingPhone:
	//	return 10010
	//case ErrTimelineStoreRequire:
	//	return 10100
	//case ErrTimelineNotFound:
	//	return 10101
	//case ErrTimelineCanNotGood:
	//	return 10102
	//case ErrAccountScoreNotEnough:
	//	return 10801
	//case ErrAccountCoinNotEnough:
	//	return 10802
	//case ErrApiRequestFail:
	//	return 20002
	case Is(err.error, AuthTokenExpired):
		return 50001
	default:
		return -1
	}
}
