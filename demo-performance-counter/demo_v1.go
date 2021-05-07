package demo_performance_counter

type App struct{}

func (a App) Run() {
	storage := NewRedisMetricsStorage()
	aggregator := NewAggregator()
	consoleReporter := NewConsoleReporter(storage, aggregator)
	consoleReporter.StartRepeatedReport(60, 60)

	emailReporter := NewEmailReporter(storage, aggregator)
	emailReporter.AddToAddress("xxx@xxx.com")
	emailReporter.StartDailyReport()

	collector := NewMetricsCollector(storage)
	collector.RecordRequest(NewRequestInfo("register", 111, 1234))
	collector.RecordRequest(NewRequestInfo("register", 222, 1234))
	collector.RecordRequest(NewRequestInfo("register", 333, 1234))
	collector.RecordRequest(NewRequestInfo("login", 444, 1234))
	collector.RecordRequest(NewRequestInfo("login", 555, 1234))
}
