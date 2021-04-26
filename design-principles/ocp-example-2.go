package design_principles

// 基于修改增加新功能

// CheckV2 第二版
func (a *Alert) CheckV2(api string, requestCount, errorCount, timeoutCount, durationOfSeconds int) {
	tps := requestCount / durationOfSeconds

	// 当接口当TPS超过某个预先设置当最大值时
	if tps > a.rule.GetMatchedRule(api).GetMaxTps() {
		a.notification.Notify(URGENCY, "这个接口的TPS超过最大阈值啦")
	}

	// 当接口请求出错数超过某个预先设置的最大值时
	if errorCount > a.rule.GetMaxErrorCount() {
		a.notification.Notify(SEVERE, "这个接口的请求错误数超过最大阈值啦")
	}

	timeoutTps := timeoutCount / durationOfSeconds

	// 当接口请求超时数超过某个预先设置的最大值时
	if timeoutTps > a.rule.GetMatchedRule(api).GetMaxTimeoutTps() {
		a.notification.Notify(URGENCY, "这个接口的请求超时数超过最大阈值啦")
	}
}
