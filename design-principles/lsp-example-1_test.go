package design_principles

import "testing"

// 修改后，TestTransporter 可以正常运行
func TestTransporter(t *testing.T) {
	// 里式替换原则
	demo := new(lspDemo)
	//demo.demoFunction(NewSecurityTransporter("", "", nil))
	demo.demoFunction(NewTransporter(nil))
}

// 修改后，TestSecurityTransporter 运行会报错 NoAuthorizationErr
func TestSecurityTransporter(t *testing.T) {
	// 里式替换原则
	demo := new(lspDemo)
	demo.demoFunction(NewSecurityTransporter("", "", nil))
}
