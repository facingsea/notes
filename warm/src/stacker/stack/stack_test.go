package stack_test

import (
	"stacker/stack"
	"testing"
)

func TestStack(t *testing.T) {
	count := 1
	var myStack stack.Stack
	assertTrue(t, myStack.Len() == 0, "expected empty stack.", count)
	count++
	assertTrue(t, myStack.Cap() == 0, "expected empty stack.", count)
}

func assertTrue(t *testing.T, condition bool, message string, id int) {
	if !condition {
		t.Errorf("#%d: %s", id, message)
	}
}

func assertEqualString(t *testing.T, a, b, string, message string, id int) {
	if a != b {
		t.Errorf("#%d: %s \"%s\" != \n\"%s\" ", id, message, a, b)
	}
}
