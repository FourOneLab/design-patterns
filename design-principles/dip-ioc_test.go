package design_principles

import (
	"fmt"
	"testing"
)

func TestDoTest(t *testing.T) {
	// 手写测试用例
	ust := new(UserServiceTest)

	if ust.DoTest() {
		fmt.Println("Test succeed.")
	} else {
		fmt.Println("Test failed.")
	}
}

func TestUserServiceTest_DoTest1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
		{
			"case-1",
			false,
		},
	}

	// 使用 Golang 提供的测试框架
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServiceTest{}
			if got := u.DoTest(); got != tt.want {
				t.Errorf("DoTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
