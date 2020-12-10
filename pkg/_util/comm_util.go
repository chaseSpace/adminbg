package _util

import (
	"fmt"
	"runtime"
)

func If(condition bool, then func(), _else ...func()) {
	if condition {
		if then != nil {
			then()
		}
	} else {
		for _, f := range _else {
			f()
		}
	}
}

func PanicIfErr(err interface{}, ignoreErrs []error, printText ...string) {
	if err != nil {
		for _, e := range ignoreErrs {
			if e == err {
				return
			}
		}
		if len(printText) > 0 {
			panic(printText[0])
		}
		panic(err)
	}
}

func Must(condition bool, err error) {
	if !condition {
		panic(err)
	}
}

// skip=1 为调用者位置，skip=2为调用者往上一层的位置，以此类推
// return-example: /develop/go/test_go/tmp_test.go:88
func FileWithLineNum(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%v:%v", file, line)
}
