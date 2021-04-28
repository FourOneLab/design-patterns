package design_principles

type Statistics struct {
	max           int64
	min           int64
	average       int64
	sum           int64
	percentile99  int64
	percentile999 int64
}

func Count(dataSet []int64) *Statistics {
	statistics := new(Statistics)

	//...省略计算逻辑...

	return statistics
}

func Max(dataSet []int64)     {}
func Min(dataSet []int64)     {}
func Average(dataSet []int64) {}
