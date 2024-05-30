package object

import "reflect"

func CallReflectMethod(any interface{}, methodName string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	if v := reflect.ValueOf(any).MethodByName(methodName); v.String() == "<invalid Value>" {
		return nil
	} else {
		return v.Call(inputs)
	}
}
