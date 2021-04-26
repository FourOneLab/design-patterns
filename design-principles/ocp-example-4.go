package design_principles

// 实现基于 Kafka 发送异步消息，将功能抽象为一组与具体消息队列无关的异步消息接口。
// 所有上层系统都依赖这组抽象的接口编程，并通过依赖注入的方式来调用。
// 可以非常方便的实现底层消息队列的替换。

type MessageQueue interface {
	Receive()
	Send(formatter MessageFormatter)
}

type KafkaMessageQueue struct{}

func (k KafkaMessageQueue) Receive() {
	panic("implement me")
}

func (k KafkaMessageQueue) Send(formatter MessageFormatter) {
	panic("implement me")
}

type RocketMQMessageQueue struct{}

func (r RocketMQMessageQueue) Receive() {
	panic("implement me")
}

func (r RocketMQMessageQueue) Send(formatter MessageFormatter) {
	panic("implement me")
}

type MessageFormatter interface {
	Format()
}

type JsonMessageFormatter struct{}

func (j JsonMessageFormatter) Format() {
	panic("implement me")
}

type ProtoBufMessageFormatter struct{}

func (p ProtoBufMessageFormatter) Format() {
	panic("implement me")
}

type Demo struct {
	MessageQueue // 基于接口而非实现编程
}

func NewDemo(messageQueue MessageQueue) *Demo { // 依赖注入
	return &Demo{MessageQueue: messageQueue}
}

func (d *Demo) SendNotification(info Infos, formatter MessageFormatter) {
	d.Send(formatter)
}

//Infos 带发送的消息
type Infos struct{}
