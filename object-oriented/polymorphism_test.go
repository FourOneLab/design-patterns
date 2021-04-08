package object_oriented

import (
	"fmt"
	"testing"
)

func TestSortedDynamicArray_Add(t *testing.T) {
	innerTest := func(a Array) {
		a.Add(5)
		a.Add(1)
		a.Add(3)
		for i := 0; i < 3; i++ {
			fmt.Print(a.Get(i))
		}
		fmt.Println()
	}

	tests := []struct {
		name   string
		fields Array
	}{
		{
			"sorted",
			NewSortedDynamicArray(),
		},
		{
			"dynamic",
			NewDynamicArray(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			innerTest(tt.fields)
		})
	}
}

func Test_prints(t *testing.T) {
	tests := []struct {
		name string
		args Iterator
	}{
		{
			"MyArray",
			MyArray{},
		},
		{
			"LinkedList",
			LinkedList{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prints(tt.args)
		})
	}
}
