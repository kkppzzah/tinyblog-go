// Package common 公用功能。
package common

import "fmt"

// Error 数据库相关错误。
type Error struct {
	Code int
	Err  error
}

// NewError 创建表示数据库相关错误的DBError。
func NewError(code int, err error) *Error {
	dbError := &Error{
		Code: code,
		Err:  err,
	}
	return dbError
}

// Error 返回Error的字符串表示。
func (err *Error) Error() string {
	if err.Err != nil {
		return fmt.Sprintf("code: %d, error: %s", err.Code, err.Err.Error())
	}
	return fmt.Sprintf("code: %d", err.Code)
}
