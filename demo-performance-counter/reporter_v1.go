package demo_performance_counter

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"time"
)

var (
	defaultSMTPServer     = ""
	defaultSenderEmail    = ""
	defaultSenderPassword = ""
)

// ConsoleReporter 和 EmailReporter 中存在代码重复问题。
// 在这两个类中，从数据库中取数据、做统计的逻辑都是相同的，可以抽取出来复用，否则就违反了 DRY 原则。
// 而且整个类负责的事情比较多，职责不是太单一。特别是显示部分的代码，可能会比较复杂（比如 Email 的展示方式），
// 最好是将显示部分的代码逻辑拆分成独立的类。除此之外，因为代码中涉及线程操作，
// 并且调用了 Aggregator 的方法，所以代码的可测试性不好。
type ConsoleReporter struct {
	metricsStorage MetricsStorage
	ticker         *time.Ticker
	aggregator     *Aggregator
}

func NewConsoleReporter(metricsStorage MetricsStorage, aggregator *Aggregator) *ConsoleReporter {
	return &ConsoleReporter{
		metricsStorage: metricsStorage,
		aggregator:     aggregator,
	}
}

func (r *ConsoleReporter) StartRepeatedReport(period, duration time.Duration) {
	r.ticker = time.NewTicker(period)

	select {
	case <-r.ticker.C:
		// 功能1：根据给定的时间区间，从数据库中获取数据
		endTime := time.Now()
		startTime := endTime.Add(-duration)
		requestInfos := r.metricsStorage.GetRequestInfos(startTime, endTime)

		stats := make(map[string]*RequestStat)
		for apiName, infos := range requestInfos {
			// 根据原始数据，计算得到统计数据
			requestStat := r.aggregator.Aggregate(infos, duration)
			stats[apiName] = requestStat
		}

		// 将统计数据显示在终端（命令行/邮件）
		fmt.Println(fmt.Sprintf("Time span: [ %s, %s ]", startTime, endTime))
		marshal, err := json.Marshal(stats)
		if err != nil {
			return
		}
		fmt.Println(string(marshal))
	}
}

type EmailReporter struct {
	metricsStorage MetricsStorage
	aggregator     *Aggregator
	emailSender    *EmailSender
	toAddress      []string
}

func NewEmailReporter(metricsStorage MetricsStorage, aggregator *Aggregator) *EmailReporter {
	return &EmailReporter{
		metricsStorage: metricsStorage,
		aggregator:     aggregator,
		emailSender:    NewEmailSender(defaultSMTPServer),
	}
}

func (r *EmailReporter) AddToAddress(addr ...string) {
	r.emailSender.AddReceiver(addr...)
}

func (r *EmailReporter) StartDailyReport() {
	duration := 24 * time.Hour
	ticker := time.NewTicker(duration)

	select {
	case <-ticker.C:
		endTime := time.Now()
		startTime := endTime.Add(-duration)
		requestInfos := r.metricsStorage.GetRequestInfos(startTime, endTime)

		stats := make(map[string]*RequestStat)
		for apiName, infos := range requestInfos {
			requestStat := r.aggregator.Aggregate(infos, duration)
			stats[apiName] = requestStat
		}
	}

	// TODO: 格式化为 HTML，并发送邮件
}

type EmailSender struct {
	smtpAddr string
	from     string
	password string
	to       []string
	auth     smtp.Auth
}

func NewEmailSender(smtpAddr string) *EmailSender {
	if smtpAddr == "" {
		smtpAddr = defaultSMTPServer
	}

	return &EmailSender{
		smtpAddr: smtpAddr,
		from:     defaultSenderEmail,
		password: defaultSenderPassword,
		auth:     smtp.PlainAuth("", defaultSenderEmail, defaultSenderPassword, smtpAddr),
	}
}

func (s *EmailSender) AddReceiver(receiver ...string) {
	s.to = append(s.to, receiver...)
}

func (s *EmailSender) SendMail(msg []byte) error {
	return smtp.SendMail(s.smtpAddr, s.auth, s.from, s.to, msg)
}
