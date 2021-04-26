package design_principles

import "sync"

// 基于开闭原则增加新功能

var (
	ApplicationContextInstance *ApplicationContext // 单例
	once                       sync.Once
)

type AlertHandler interface {
	Check(info *ApiStatInfo)
}

type ApiStatInfo struct {
	api               string
	requestCount      int
	errorCount        int
	durationOfSeconds int
	timeoutCount      int // feat：接口超时检查
}

func NewDefaultApiStatInfo() *ApiStatInfo {
	return &ApiStatInfo{
		api:               "/",
		requestCount:      20,
		errorCount:        20,
		durationOfSeconds: 20,
		timeoutCount:      20,
	}
}

func (i *ApiStatInfo) GetApi() string {
	return i.api
}

func (i *ApiStatInfo) GetRequestCount() int {
	return i.requestCount
}

func (i *ApiStatInfo) GetErrorCount() int {
	return i.errorCount
}

func (i *ApiStatInfo) GetDurationOfSeconds() int {
	return i.durationOfSeconds
}

func (i *ApiStatInfo) GetTimeoutCount() int {
	return i.timeoutCount
}

type AlertV3 struct {
	alertHandlers []AlertHandler
}

func NewAlertV3() *AlertV3 {
	return &AlertV3{}
}

func (a *AlertV3) AddAlertHandler(handler ...AlertHandler) {
	a.alertHandlers = append(a.alertHandlers, handler...)
}

func (a *AlertV3) Check(info *ApiStatInfo) {
	for _, handler := range a.alertHandlers {
		handler.Check(info)
	}
}

type BaseHandler struct {
	rule         *AlertRuleV3
	notification *NotificationV3
}

func NewBaseHandler(rule *AlertRuleV3, notification *NotificationV3) *BaseHandler {
	return &BaseHandler{
		rule:         rule,
		notification: notification,
	}
}

type AlertRuleV3 struct {
	maxTps        int
	maxErrorCount int
	maxTimeoutTps int
}

func NewAlertRuleV3(maxTps int, maxErrorCount int, maxTimeoutTps int) *AlertRuleV3 {
	return &AlertRuleV3{
		maxTps:        maxTps,
		maxErrorCount: maxErrorCount,
		maxTimeoutTps: maxTimeoutTps,
	}
}

func (a *AlertRuleV3) GetMatchedRule(api string) *AlertRuleV3 {
	return a
}

func (a *AlertRuleV3) GetMaxTps() int {
	return a.maxTps
}

func (a *AlertRuleV3) GetMaxErrorCount() int {
	return a.maxErrorCount
}

func (a *AlertRuleV3) GetMaxTimeoutTps() int {
	return a.maxTimeoutTps
}

type NotificationV3 struct{}

func NewNotificationV3() *NotificationV3 {
	return &NotificationV3{}
}

func (n NotificationV3) Notify(level NotificationEmergencyLevel, info string) {}

type TpsAlertHandler struct {
	*BaseHandler
}

func NewTpsAlertHandler(rule *AlertRuleV3, notification *NotificationV3) *TpsAlertHandler {
	return &TpsAlertHandler{BaseHandler: NewBaseHandler(rule, notification)}
}

func (t *TpsAlertHandler) Check(info *ApiStatInfo) {
	tps := info.GetRequestCount() / info.GetDurationOfSeconds()

	if tps > t.rule.GetMatchedRule(info.GetApi()).GetMaxTps() {
		t.notification.Notify(URGENCY, "这个接口的TPS超过最大阈值啦")
	}
}

type ErrorAlertHandler struct {
	*BaseHandler
}

func NewErrorAlertHandler(rule *AlertRuleV3, notification *NotificationV3) *ErrorAlertHandler {
	return &ErrorAlertHandler{BaseHandler: NewBaseHandler(rule, notification)}
}

func (e *ErrorAlertHandler) Check(info *ApiStatInfo) {
	if info.GetErrorCount() > e.rule.GetMatchedRule(info.GetApi()).GetMaxErrorCount() {
		e.notification.Notify(SEVERE, "这个接口的请求错误数超过最大阈值啦")
	}
}

type TimeoutAlertHandler struct {
	*BaseHandler
}

func NewTimeoutAlertHandler(rule *AlertRuleV3, notification *NotificationV3) *TimeoutAlertHandler {
	return &TimeoutAlertHandler{BaseHandler: NewBaseHandler(rule, notification)}
}

func (t *TimeoutAlertHandler) Check(info *ApiStatInfo) {
	if info.timeoutCount > t.rule.GetMaxTimeoutTps() {
		t.notification.Notify(URGENCY, "这个接口的请求超时数超过最大阈值啦")
	}
}

// ApplicationContext 负责 Alert 的创建、组装（alertRule 和 notification 的依赖注入）、初始化（添加 handlers）工作
type ApplicationContext struct {
	alertRule    *AlertRuleV3
	notification *NotificationV3
	alert        *AlertV3
}

func NewApplicationContext() *ApplicationContext {
	alertRule := NewAlertRuleV3(10, 10, 10)
	notification := NewNotificationV3()

	alert := NewAlertV3()
	alert.AddAlertHandler(
		NewTpsAlertHandler(alertRule, notification),
		NewErrorAlertHandler(alertRule, notification),
		NewTimeoutAlertHandler(alertRule, notification))

	return &ApplicationContext{
		alertRule:    alertRule,
		notification: notification,
		alert:        alert,
	}
}

func (a *ApplicationContext) GetAlert() *AlertV3 {
	return a.alert
}

func GetApplicationContextInstance() *ApplicationContext {
	once.Do(func() {
		ApplicationContextInstance = NewApplicationContext()
	})

	return ApplicationContextInstance
}
