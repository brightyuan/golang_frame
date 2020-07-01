package _package

import "testing"

/**
测试package 方法 CtoF
 */
func TestCtoF(t *testing.T) {
	tt := CtoF(54)
	expect := 86.0

	if tt != expect {
		t.Errorf("%v ,actual %v", tt, expect)
	}
}
