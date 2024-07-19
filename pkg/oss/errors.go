package oss

import (
	"errors"
	"fmt"
)

var ParamError = errors.New("Oss parameters error")
var AuthError = errors.New("Oss authorization error")
var EncryptError = errors.New("Oss encrypted error")
var PathError = errors.New("Oss path error")

func WrapOssError(errType error, message string) error {
	return fmt.Errorf("%w: %s", errType, message)
}
