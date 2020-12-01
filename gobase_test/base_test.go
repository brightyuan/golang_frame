package gobase_test

import "go/gobase"
import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var c = gobase.Add(1, 2)
	fmt.Println(c)

}
