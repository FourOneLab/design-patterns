package design_principles

// ------------------------------
// 非依赖注入实现

type NotificationNoDI struct {
	*MessageSender
}

func NewNotificationNoDI() *NotificationNoDI {
	return &NotificationNoDI{
		MessageSender: NewMessageSender(), // 此处有点像 hardcode
	}
}

func (n *NotificationNoDI) SendMessage(cellphone, message string) {
	n.MessageSender.Send(cellphone, message) // 这种写法，当 NotificationNoDI 重载 Send 方法时不用修改
}

type MessageSender struct{}

func NewMessageSender() *MessageSender {
	return &MessageSender{}
}

func (m MessageSender) Send(cellphone, message string) {}

// ------------------------------
// 依赖注入实现

type NotificationDI struct {
	*MessageSender
}

func NewNotificationDI(messageSender *MessageSender) *NotificationDI {
	return &NotificationDI{
		MessageSender: messageSender, // 通过 new 函数将 messageSender 传递进来
	}
}

func (n *NotificationDI) Notification(cellphone, message string) {
	n.MessageSender.Send(cellphone, message)
}

// ------------------------------
// 依赖注入并基于接口而非实现编程

type MsgSender interface {
	Send(cellphone, message string)
}

type NotificationDIV2 struct {
	MsgSender
}

func NewNotificationDIV2(msgSender MsgSender) *NotificationDIV2 {
	return &NotificationDIV2{MsgSender: msgSender}
}

func (n *NotificationDIV2) Notification(cellphone, message string) {
	n.MsgSender.Send(cellphone, message)
}

// SMSSender 短信发送类
type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(cellphone, message string) {
	panic("implement me")
}

// InboxSender 站内信发送类
type InboxSender struct{}

func NewInboxSender() *InboxSender {
	return &InboxSender{}
}

func (i *InboxSender) Send(cellphone, message string) {
	panic("implement me")
}
