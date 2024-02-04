package util

import "fmt"

func Fail(err error, msg string) (int64, error) {
	return 0, fmt.Errorf("%s: %v", msg, err)
}