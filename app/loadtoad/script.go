package loadtoad

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/kyleu/loadtoad/app/lib/har"
	"github.com/kyleu/loadtoad/app/util"
	"github.com/pkg/errors"
)

func scriptUpdateEntry(vm *goja.Runtime, e *har.Entry) error {
	if f, ok := goja.AssertFunction(vm.Get("updateRequest")); ok {
		reqVM := vm.ToValue(e.Request)
		ret, err := f(goja.Undefined(), reqVM)
		if err != nil {
			return err
		}
		_, ok := ret.Export().(*har.Request)
		if !ok {
			return errors.Errorf("return value of [%s] is [%T], not [%s]", "updateRequest", ret.Export(), "*har.Request")
		}
	}
	if f, ok := goja.AssertFunction(vm.Get("updateResponse")); ok {
		rspVM := vm.ToValue(e.Response)
		ret, err := f(goja.Undefined(), rspVM)
		if err != nil {
			return err
		}
		_, ok := ret.Export().(*har.Response)
		if !ok {
			return errors.Errorf("return value of [%s] is [%T], not [%s]", "updateResponse", ret.Export(), "*har.Response")
		}
	}
	return nil
}

func scriptExtractVariables(vm *goja.Runtime, wr *WorkflowResult) (util.ValueMap, error) {
	ret := util.ValueMap{}
	if f, ok := goja.AssertFunction(vm.Get("extractVariables")); ok {
		reqVM, rspVM := vm.ToValue(wr.Entry.Request), vm.ToValue(wr.Response)
		jsRet, err := f(goja.Undefined(), reqVM, rspVM)
		if err != nil {
			return nil, err
		}
		vars, ok := jsRet.Export().(map[string]any)
		if !ok {
			return nil, errors.Errorf("return value of [%s] is [%T], not [%s]", "extractVariables", jsRet.Export(), "map[string]any")
		}
		if len(vars) > 0 {
			ret = ret.Merge(vars)
		}
	}
	return ret, nil
}

func scriptExtractReplacements(vm *goja.Runtime, e *har.Entry) (map[string]string, error) {
	ret := map[string]string{}
	if f, ok := goja.AssertFunction(vm.Get("extractReplacements")); ok {
		reqVM, rspVM := vm.ToValue(e.Request), vm.ToValue(e.Response)
		jsRet, err := f(goja.Undefined(), reqVM, rspVM)
		if err != nil {
			return nil, err
		}
		repls, ok := jsRet.Export().(map[string]any)
		if !ok {
			return nil, errors.Errorf("return value of [%s] is [%T], not [%s]", "extractReplacements", jsRet.Export(), "map[string]any")
		}
		for k, v := range repls {
			ret[k] = fmt.Sprint(v)
		}
	}
	return ret, nil
}
