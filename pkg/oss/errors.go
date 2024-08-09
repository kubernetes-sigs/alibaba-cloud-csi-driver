package oss

import (
	"errors"
	"fmt"
)

var ParamError = errors.New("OSS parameters error")
var AuthError = errors.New("OSS authorization error")
var EncryptError = errors.New("OSS encrypted error")
var PathError = errors.New("OSS path error")

func WrapOssError(errType error, format string, args ...interface{}) error {
	return fmt.Errorf("%w: %s", errType, fmt.Sprintf(format, args...))
}
