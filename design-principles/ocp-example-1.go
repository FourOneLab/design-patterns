package design_principles

// API 接口监控告警
const (
	SEVERE  NotificationEmergencyLevel = iota // 严重
	URGENCY                                   // 紧急
	NORMAL                                    // 普通
	TRIVIAL                                   // 无关紧要
)

type NotificationEmergencyLevel int

// AlertRule 存储告警规则
type AlertRule struct {
	maxTps        int
	maxErrorCount int
	maxTimeoutTps int
}

func (a *AlertRule) GetMatchedRule(api string) *AlertRule {
	return a
}

func (a *AlertRule) GetMaxTps() int {
	return a.maxTps
}

func (a *AlertRule) GetMaxErrorCount() int {
	return a.maxErrorCount
}

func (a *AlertRule) GetMaxTimeoutTps() int {
	return a.maxTimeoutTps
}

// Notification 告警通知类，支持多种通知渠道：邮件，短信，微信，手机
type Notification struct{}

func (n *Notification) Notify(level NotificationEmergencyLevel, info string) {

}

type Alert struct {
	rule         AlertRule
	notification Notification
}

func NewAlert(rule AlertRule, notification Notification) *Alert {
	return &Alert{
		rule:         rule,
		notification: notification,
	}
}

// CheckV1 第一版
func (a *Alert) CheckV1(api string, requestCount, errorCount, durationOfSeconds int) {
	tps := requestCount / durationOfSeconds

	// 当接口当TPS超过某个预先设置当最大值时
	if tps > a.rule.GetMatchedRule(api).GetMaxTps() {
		a.notification.Notify(URGENCY, "...")
	}

	// 当接口请求出错数超过某个预先设置的最大值时
	if errorCount > a.rule.GetMaxErrorCount() {
		a.notification.Notify(SEVERE, "...")
	}
}
