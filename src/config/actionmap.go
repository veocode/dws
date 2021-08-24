package config

import (
	"reflect"

	"github.com/veocode/dws/src/actions"
	"github.com/veocode/dws/src/contracts"
)

var defaultAction = actions.Version{}

var actionMap map[string]reflect.Type = map[string]reflect.Type{
	"":        reflect.TypeOf(defaultAction),
	"test":    reflect.TypeOf(actions.Test{}),
	"version": reflect.TypeOf(actions.Version{}),
	"init":    reflect.TypeOf(actions.Init{}),
}

type ActionMapper struct {
}

func (am *ActionMapper) IsActionExists(actionName string) bool {
	return actionMap[actionName] != nil
}

func (am *ActionMapper) GetActionHandler(actionName string) contracts.Action {
	actionHandler := reflect.New(actionMap[actionName]).Elem()
	return actionHandler.Interface().(contracts.Action)
}
