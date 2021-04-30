package design_principles

import "testing"

func TestName(t *testing.T) {
	// 非依赖注入方式
	NewNotificationNoDI()

	// 依赖注入方式
	messageSender := NewMessageSender()    // 创建对象
	di := NewNotificationDI(messageSender) //依赖注入
	di.Notification("12345678900", "这是要发送的消息")

	// 依赖注入并基于接口而非实现编程
	smsSender := NewSMSSender()            // 创建对象
	div2 := NewNotificationDIV2(smsSender) // 依赖注入
	div2.Notification("12345678900", "短信验证码：123456")
}
