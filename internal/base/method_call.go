package base

import (
	"errors"
	"fmt"
	"gengine/context"
	"reflect"
	"runtime"
	"strings"
)

type MethodCall struct {
	SourceCode
	MethodName string
	MethodArgs *Args
	dataCtx    *context.DataContext
}

func (mc *MethodCall) Initialize(dataCtx *context.DataContext) {
	mc.dataCtx = dataCtx

	if mc.MethodArgs != nil {
		mc.MethodArgs.Initialize(dataCtx)
	}
}

func (mc *MethodCall) AcceptArgs(funcArg *Args) error {
	if mc.MethodArgs == nil {
		mc.MethodArgs = funcArg
		return nil
	}
	return errors.New("methodArgs set twice!")
}

func (mc *MethodCall) Evaluate(Vars map[string]reflect.Value) (mr reflect.Value, err error) {

	defer func() {
		if e := recover(); e != nil {
			size := 1 << 10 * 10
			buf := make([]byte, size)
			rs := runtime.Stack(buf, false)
			if rs > size {
				rs = size
			}
			buf = buf[:rs]
			eMsg := fmt.Sprintf("line %d, column %d, code: %s, %+v \n%s", mc.LineNum, mc.Column, mc.Code, e, string(buf))
			eMsg = strings.ReplaceAll(eMsg, "panic", "error")
			err = errors.New(eMsg)
		}
	}()

	var argumentValues []reflect.Value
	if mc.MethodArgs == nil {
		argumentValues = make([]reflect.Value, 0)
	} else {
		av, err := mc.MethodArgs.Evaluate(Vars)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		argumentValues = av
	}

	mr, err = mc.dataCtx.ExecMethod(Vars, mc.MethodName, argumentValues)
	if err != nil {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("line %d, column %d, code: %s, %+v", mc.LineNum, mc.Column, mc.Code, err))
	}
	return
}
