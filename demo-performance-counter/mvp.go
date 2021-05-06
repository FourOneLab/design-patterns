package demo_performance_counter

import (
	"encoding/json"
	"fmt"
	"time"
)

type Metric interface {
	RecordResponseTime(apiName string, responseTime time.Duration)
	RecordTimestamp(apiName string, timestamp int64)
	StartRepeatedReport(period time.Duration) error
}

type Metrics struct {
	responseTimes map[string][]time.Duration // key 是接口名称，value 是对应接口请求的响应时间
	timestamps    map[string][]int64         // key 是接口名称，value 是对应接口请求的时间戳
	ticker        *time.Ticker
}

// RecordResponseTime 记录接口请求的响应时间
func (m *Metrics) RecordResponseTime(apiName string, responseTime time.Duration) {
	if _, ok := m.responseTimes[apiName]; !ok {
		m.responseTimes[apiName] = make([]time.Duration, 0)
	}

	m.responseTimes[apiName] = append(m.responseTimes[apiName], responseTime)
}

// RecordTimestamp 记录接口请求的访问时间
func (m *Metrics) RecordTimestamp(apiName string, timestamp int64) {
	if _, ok := m.timestamps[apiName]; !ok {
		m.timestamps[apiName] = make([]int64, 0)
	}

	m.timestamps[apiName] = append(m.timestamps[apiName], timestamp)
}

func (m *Metrics) StartRepeatedReport(period time.Duration) error {
	m.ticker = time.NewTicker(period)
	stats := make(map[string]map[string]int64)

	for apiName, durations := range m.responseTimes {
		if _, ok := stats[apiName]; !ok {
			stats[apiName] = make(map[string]int64)
		}

		stats[apiName]["max"] = int64(maxTimeDuration(durations))
		stats[apiName]["avg"] = int64(avgTimeDuration(durations))
	}

	for apiName, timestamps := range m.timestamps {
		if _, ok := stats[apiName]; !ok {
			stats[apiName] = make(map[string]int64)
		}

		stats[apiName]["count"] = int64(len(timestamps))
	}

	marshal, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	fmt.Println(string(marshal))

	return nil
}

func maxTimeDuration(dataset []time.Duration) time.Duration {
	if len(dataset) < 1 {
		return time.Duration(0)
	}

	max := dataset[0]

	for i, duration := range dataset {
		if max < duration {
			max = dataset[i]
		}
	}

	return max
}

func avgTimeDuration(dataset []time.Duration) time.Duration {
	if len(dataset) < 1 {
		return time.Duration(0)
	}

	sum := time.Duration(0)

	for _, duration := range dataset {
		sum += duration
	}

	res := int64(sum) / int64(len(dataset))

	return time.Duration(res)
}
