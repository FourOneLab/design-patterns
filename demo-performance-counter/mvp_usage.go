package demo_performance_counter

import (
	"log"
	"time"
)

// 应用场景：统计下面2个接口（注册/登录）的响应时间和访问次数

type UserVO struct{}

type UserController struct {
	Metric
}

func NewUserController(metric Metric) *UserController {
	if err := metric.StartRepeatedReport(60 * time.Second); err != nil {
		log.Println(err)
		return nil
	}

	return &UserController{Metric: metric}
}

func (c *UserController) Register(user UserVO) {
	startTimestamp := time.Now().Unix()
	c.RecordTimestamp("register", startTimestamp)

	respTime := time.Now().Unix() - startTimestamp
	c.RecordResponseTime("register", time.Duration(respTime))
}

func (c *UserController) Login(telephone, password string) {
	startTimestamp := time.Now().Unix()
	c.RecordTimestamp("login", startTimestamp)

	respTime := time.Now().Unix() - startTimestamp
	c.RecordResponseTime("login", time.Duration(respTime))
}
