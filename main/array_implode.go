package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	//v := Implode(", ", 1, "2", "0.2", .1, S{});
	var data = []string{"dsf","sdfas","dasfas"}
	//var data = []float32{1, 2, 3, 4, 5, 6, 6}
	fmt.Println(Implode(data, ","))
}

func Implode(list interface{}, seq string) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}
	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}

func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}
