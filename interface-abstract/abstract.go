package interface_abstract

import (
	"io"
	"os"

	"github.com/Shopify/sarama"
)

// 抽象类的一个典型使用场景（模板设计模式）

// Golang中没有抽象类这个语法特性，所以使用组合来实现。

type Level struct {
	Value int
}

func (l Level) intValue() int {
	return l.Value
}

// logger 是一个记录日志的抽象类，FileLogger和MessageQueueLogger继承logger。
// 模拟抽象类，所以不导出
type logger struct {
	name              string
	enabled           bool
	minPermittedLevel Level
}

func newLogger(name string, enabled bool, minPermittedLevel Level) *logger {
	return &logger{
		name:              name,
		enabled:           enabled,
		minPermittedLevel: minPermittedLevel,
	}
}

func (l logger) isLoggable(level Level) bool {
	return l.enabled && (l.minPermittedLevel.intValue() <= level.intValue())
}

// doLog 抽象方法
func (l logger) doLog(level Level, message string) {}

type FileLogger struct {
	*logger
	writer io.Writer
}

func NewFileLogger(name string, enabled bool, minPermittedLevel Level, filePath string) (*FileLogger, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileLogger := &FileLogger{
		logger: newLogger(name, enabled, minPermittedLevel),
		writer: file,
	}

	return fileLogger, nil
}

func (l FileLogger) Log(level Level, message string) {
	if l.isLoggable(level) {
		l.doLog(level, message)
	}
}

func (l FileLogger) doLog(level Level, message string) {
	l.writer.Write([]byte(message))
}

type MessageQueueLogger struct {
	*logger
	messageQueueClient sarama.AsyncProducer
}

func NewMessageQueueLogger(name string, enabled bool, minPermittedLevel Level, msgQueueProducer sarama.AsyncProducer) *MessageQueueLogger {
	messageQueueLogger := &MessageQueueLogger{
		logger:             newLogger(name, enabled, minPermittedLevel),
		messageQueueClient: msgQueueProducer,
	}
	return messageQueueLogger
}

func (l MessageQueueLogger) Log(level Level, message string) {
	if l.isLoggable(level) {
		l.doLog(level, message)
	}
}

func (l MessageQueueLogger) doLog(level Level, message string) {
	l.messageQueueClient.Input() <- &sarama.ProducerMessage{
		Topic: "logger",
		Value: sarama.StringEncoder(message),
	}
}
