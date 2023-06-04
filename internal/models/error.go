package models

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var ErrRecordNotFound = errors.New("Record Not Found")

// DetailError will wrap the err with detail runtime message,
// such as filename, line number and the method name
func DetailError(err error) error {
	pc, filename, line, _ := runtime.Caller(1)
	methodName := runtime.FuncForPC(pc).Name()
	m := strings.Split(methodName, ".")
	methodName = m[len(m)-1]
	return fmt.Errorf("%s line:%v %s(): %w", filename, line, methodName, err)
}
