package utils

import (
	"fmt"
	"testing"
)

func TestCheckDateTime(t *testing.T) {
	dt := "07:45"
	isValid := CheckDateTime(dt)
	expect := true
	if isValid != expect {
		t.Errorf("CheckDateTime(%s) = %v; expected %v", dt, isValid, expect)
	}
	fmt.Println("ok")
}
