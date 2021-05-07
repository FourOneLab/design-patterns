package demo_performance_counter

import (
	"time"
)

// MetricsCollector 负责采集和存储数据，职责相对来说比较单一。
// 它基于接口而非实现编程，通过依赖注入的方式来传递 MetricsStorage 对象，
// 可以在不需要修改代码的情况下，灵活地替换不同的存储方式，满足开闭原则。
type MetricsCollector struct {
	metricsStorage MetricsStorage // 基于接口而非实现编程
}

// NewMetricsCollector 依赖注入的方式新建 collector
func NewMetricsCollector(metricsStorage MetricsStorage) *MetricsCollector {
	return &MetricsCollector{metricsStorage: metricsStorage}
}

// RecordRequest 用一个方法替代 MVP 中的两个方法
func (c *MetricsCollector) RecordRequest(info *RequestInfo) {
	if info == nil || info.ApiName() == "" {
		return
	}
	c.metricsStorage.SaveRequestInfo(info)
}

type RequestInfo struct {
	apiName      string
	responseTime time.Duration
	timestamp    int64
}

func NewRequestInfo(apiName string, responseTime time.Duration, timestamp int64) *RequestInfo {
	return &RequestInfo{
		apiName:      apiName,
		responseTime: responseTime,
		timestamp:    timestamp,
	}
}

func (r *RequestInfo) ApiName() string {
	return r.apiName
}

func (r *RequestInfo) SetApiName(apiName string) {
	r.apiName = apiName
}

func (r *RequestInfo) ResponseTime() time.Duration {
	return r.responseTime
}

func (r *RequestInfo) SetResponseTime(responseTime time.Duration) {
	r.responseTime = responseTime
}

func (r *RequestInfo) Timestamp() int64 {
	return r.timestamp
}

func (r *RequestInfo) SetTimestamp(timestamp int64) {
	r.timestamp = timestamp
}
