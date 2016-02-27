package stack_test

import (
	"fmt"
	"stacker/stack"
	"testing"
)

func TestStack(t *testing.T) {
	count := 1
	var myStack stack.Stack
	assertTrue(t, myStack.Len() == 0, "expected empty stack.", count) //1
	count++
	assertTrue(t, myStack.Cap() == 0, "expected empty stack.", count) //2
	count++
	assertTrue(t, myStack.IsEmpty(), "expected empty stack.", count) //3
	count++
	value, err := myStack.Pop()
	assertTrue(t, value == nil, "expected nil value", count) //4
	count++
	assertTrue(t, err != nil, "expected error", count) //5
	count++
	value1, err := myStack.Top()
	assertTrue(t, value1 == nil, "expected nil value", count) //6
	count++
	assertTrue(t, err != nil, "expected error", count) //7
	count++
	myStack.Push(1)
	myStack.Push(12.34)
	myStack.Push("hello world")
	assertTrue(t, myStack.Len() == 3, "expected nonempey stack.", count) //8
	count++
	assertTrue(t, myStack.IsEmpty() == false, "expected nonempty stack.", count) //9
	count++
	value2, err := myStack.Pop()
	//value2为interface{}类型，需要转换为string类型
	assertEqualString(t, value2.(string), "hello world", "unexpected text.", count) //10
	count++
	assertTrue(t, err == nil, "no expected error.", count) //11
	count++
	value3, err := myStack.Top()
	count++
	fmt.Println(value3, count)
	assertTrue(t, value3 == 12.34, "unexpected value.", count) //12
	count++
	assertTrue(t, err == nil, "no error expected.", count) //13
	count++
	myStack.Pop()
	assertTrue(t, myStack.Len() == 1, "expected nonempty stack.", count) //14
	count++
	assertTrue(t, myStack.IsEmpty() == false, "expected nonempty stack.", count) //15
	count++
	value4, err := myStack.Pop()
	assertTrue(t, value4 == 1, "unexpected number.", count) //16
	count++
	assertTrue(t, err == nil, "no error expected.", count) //17
	count++
	assertTrue(t, myStack.Len() == 0, "expected empty stack.", count) //18
	count++
	assertTrue(t, myStack.IsEmpty(), "expected empty stack.", count) //19
	count++
}

func assertTrue(t *testing.T, condition bool, message string, id int) {
	if !condition {
		t.Errorf("#%d: %s", id, message)
	}
}

func assertEqualString(t *testing.T, a, b string, message string, id int) {
	if a != b {
		t.Errorf("#%d: %s \"%s\" != \n\"%s\" ", id, message, a, b)
	}
}
