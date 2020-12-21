package base

import (
	"gengine/context"
	"gengine/internal/core/errors"
	"reflect"
)

type ExpressionAtom struct {
	SourceCode
	Variable     string
	Constant     *Constant
	FunctionCall *FunctionCall
	MethodCall   *MethodCall
	MapVar       *MapVar
	dataCtx      *context.DataContext
}

func (e *ExpressionAtom) Initialize(dc *context.DataContext) {
	e.dataCtx = dc

	if e.Constant != nil {
		e.Constant.Initialize(dc)
	}

	if e.FunctionCall != nil {
		e.FunctionCall.Initialize(dc)
	}

	if e.MethodCall != nil {
		e.MethodCall.Initialize(dc)
	}

	if e.MapVar != nil {
		e.MapVar.Initialize(dc)
	}
}

func (e *ExpressionAtom) AcceptVariable(name string) error {
	if len(e.Variable) == 0 {
		e.Variable = name
		return nil
	}
	return errors.New("Variable already defined")
}

func (e *ExpressionAtom) AcceptConstant(cons *Constant) error {
	if e.Constant == nil {
		e.Constant = cons
		return nil
	}
	return errors.New("Constant already defined")
}

func (e *ExpressionAtom) AcceptFunctionCall(funcCall *FunctionCall) error {
	if e.FunctionCall == nil {
		e.FunctionCall = funcCall
		return nil
	}
	return errors.New("FunctionCall already defined")
}

func (e *ExpressionAtom) AcceptMethodCall(methodCall *MethodCall) error {
	if e.MethodCall == nil {
		e.MethodCall = methodCall
		return nil
	}
	return errors.New("MethodCall already defined")
}

func (e *ExpressionAtom) AcceptMapVar(mapVar *MapVar) error {
	if e.MapVar == nil {
		e.MapVar = mapVar
		return nil
	}
	return errors.New("MapVar already defined")
}

func (e *ExpressionAtom) Evaluate(Vars map[string]reflect.Value) (reflect.Value, error) {
	if len(e.Variable) > 0 {
		return e.dataCtx.GetValue(Vars, e.Variable)
	} else if e.Constant != nil {
		return e.Constant.Evaluate(Vars)
	} else if e.FunctionCall != nil {
		return e.FunctionCall.Evaluate(Vars)
	} else if e.MethodCall != nil {
		return e.MethodCall.Evaluate(Vars)
	} else if e.MapVar != nil {
		return e.MapVar.Evaluate(Vars)
	}
	//todo
	return reflect.ValueOf(nil), errors.New("ExpressionAtom Evaluate error!")
}
