package gobase_test

import (
	"ext/gobase"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var c = gobase.Add(1, 2)
	fmt.Println(c)

}
