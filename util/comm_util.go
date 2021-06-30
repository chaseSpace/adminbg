package util

import (
	"adminbg/cerror"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const (
	EnvTest = "test"
	EnvDev  = "dev"
	EnvProd = "prod"
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

func CheckSplitPageParams(pn, ps uint16) error {
	if pn == 0 || ps == 0 || ps > 100 {
		return cerror.ErrInvalidSplitPageParams
	}
	return nil
}

// split args pass with `.` or `/`
func GetFuncName(funcObj interface{}, split ...string) string {
	fn := runtime.FuncForPC(reflect.ValueOf(funcObj).Pointer()).Name()
	if len(split) > 0 {
		fs := strings.Split(fn, split[0])
		return fs[len(fs)-1]
	}
	return fn
}

// split args pass with `.` or `/`
func GetRunningFuncName(split ...string) string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	fn := runtime.FuncForPC(pc[0]).Name()

	if len(split) > 0 {
		fs := strings.Split(fn, split[0])
		return fs[len(fs)-1]
	}
	return fn
}
