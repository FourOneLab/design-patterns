package demo_performance_counter

import "time"

// MetricsStorage 和 RedisMetricsStorage 的设计比较简单。
// 需要实现新的存储方式的时候，只需要实现 MetricsStorage 接口即可。
// 因为所有用到 MetricsStorage 和 RedisMetricsStorage 的地方，
// 都是基于相同的接口函数来编程的，所以，
// 除了在组装类的地方有所改动（从 RedisMetricsStorage 改为新的存储实现类），
// 其他接口函数调用的地方都不需要改动，满足开闭原则。
type MetricsStorage interface {
	SaveRequestInfo(info *RequestInfo)

	// 注意下面两个方法！
	// 一次性取太长时间区间的数据，可能会导致拉取太多的数据到内存中，有可能会撑爆内存。
	// 有可能会触发 OOM（Out Of Memory），即便不出现 OOM 内存还够用，但也会因为内存吃紧，
	// 导致频繁的 GC，进而导致系统接口请求处理变慢，甚至超时。

	GetRequestInfo(apiName string, startTime, endTime time.Time) []RequestInfo
	GetRequestInfos(startTime, endTime time.Time) map[string][]RequestInfo
}

type RedisMetricsStorage struct{}

func NewRedisMetricsStorage() *RedisMetricsStorage {
	return &RedisMetricsStorage{}
}

func (r *RedisMetricsStorage) SaveRequestInfo(info *RequestInfo) {
	panic("implement me")
}

func (r *RedisMetricsStorage) GetRequestInfo(apiName string, startTime, endTime time.Time) []RequestInfo {
	panic("implement me")
}

func (r *RedisMetricsStorage) GetRequestInfos(startTime, endTime time.Time) map[string][]RequestInfo {
	panic("implement me")
}
