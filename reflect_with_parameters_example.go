package main

import (
	"fmt"
	"reflect"
)

type Runtime struct {
	Name string
}

func (runtime Runtime) ThisIsTheFunction(param string) (test string) {
	return param
}

func test(function_name string, function_param string) (output string) {

	runtime := Runtime{function_name}
	strucReflected := reflect.ValueOf(runtime)
	method := strucReflected.MethodByName(function_name)

	params := []reflect.Value{reflect.ValueOf(function_param)}
	temp_result := method.Call(params)
	result := temp_result[0].Interface().(string)

	return result
}

func main() {
	output := test("ThisIsTheFunction", "Param_XXX")
	fmt.Println(output)
}
