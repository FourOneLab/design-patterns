package demo_performance_counter

import (
	"math"
	"sort"
	"time"
)

// Aggregator 是一个工具类，目前只有一个方法，几十行左右的代码量，负责各种统计数据的计算。
// 当需要扩展新的统计功能时，要修改 aggregate() 方法的代码，并且一旦越来越多的统计功能添加进来之后，
// 这个函数的代码量会持续增加，可读性、可维护性就变差了。
// 这个类的设计存在职责不够单一、不易扩展等问题，需要在之后的版本中，对其结构做优化。
type Aggregator struct{}

func NewAggregator() *Aggregator {
	return &Aggregator{}
}

// Aggregate 根据原始数据，计算得到统计数据
func (a *Aggregator) Aggregate(requestInfos []RequestInfo, duration time.Duration) *RequestStat {
	maxRespTime := time.Duration(math.MinInt64)
	minRespTime := time.Duration(math.MaxInt64)
	avgRespTime := time.Duration(-1)
	p999RespTime := time.Duration(-1)
	p99RespTime := time.Duration(-1)
	sumRespTime := time.Duration(0)
	count := int64(0)

	for _, info := range requestInfos {
		count++
		respTime := info.ResponseTime()

		if maxRespTime < respTime {
			maxRespTime = respTime
		}

		if minRespTime > respTime {
			minRespTime = respTime
		}

		sumRespTime += respTime
	}

	if count != 0 {
		avgRespTime = time.Duration(int64(sumRespTime) / count)
	}

	tps := count / int64(duration)

	sort.Slice(requestInfos, func(i, j int) bool {
		if requestInfos[i].ResponseTime() >= requestInfos[j].ResponseTime() {
			return true
		}
		return false
	})

	idx999 := float64(count) * 0.999
	idx99 := float64(count) * 0.99

	if count != 0 {
		p999RespTime = requestInfos[int(idx999)].ResponseTime()
		p99RespTime = requestInfos[int(idx99)].ResponseTime()
	}

	stat := &RequestStat{
		MaxResponseTime:  maxRespTime,
		MinResponseTime:  minRespTime,
		AvgResponseTime:  avgRespTime,
		P999ResponseTime: p999RespTime,
		P99ResponseTime:  p99RespTime,
		Count:            count,
		Tps:              tps,
	}

	return stat
}

type RequestStat struct {
	MaxResponseTime  time.Duration
	MinResponseTime  time.Duration
	AvgResponseTime  time.Duration
	P999ResponseTime time.Duration
	P99ResponseTime  time.Duration
	Count            int64
	Tps              int64
}
