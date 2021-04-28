package design_principles

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Viewer interface {
	OutputInPlainText() string
	Output() map[string]string
}

type Updater interface {
	update() // 从 ConfigSource 加载配置到 address/timeout/maxTotal...
}

// ConfigSource 配置中心，如 Zookeeper/Apollo
type ConfigSource interface {
	config()
}

type ZookeeperConfigSource struct{}

func NewZookeeperConfigSource() *ZookeeperConfigSource {
	return &ZookeeperConfigSource{}
}

func (z *ZookeeperConfigSource) config() {
	log.Println("read config from zookeeper")
}

type ApolloConfigSource struct{}

func NewApolloConfigSource() *ApolloConfigSource {
	return &ApolloConfigSource{}
}

func (a *ApolloConfigSource) config() {
	log.Println("read config from apollo")
}

// ------------------------------

// RedisConfig 配置类
type RedisConfig struct {
	ConfigSource
	address       string
	timeout       time.Duration
	maxTotal      int
	maxWaitMillis time.Duration
	maxIdle       int
	minIdle       int
}

func NewRedisConfig(configSource ConfigSource) *RedisConfig {
	return &RedisConfig{ConfigSource: configSource}
}

func (r *RedisConfig) OutputInPlainText() string {
	panic("implement me")
}

func (r *RedisConfig) Output() map[string]string {
	panic("implement me")
}

func (r RedisConfig) Address() string {
	return r.address
}

func (r *RedisConfig) update() {
	r.config()
}

// ------------------------------

// KafkaConfig Kafka配置类
type KafkaConfig struct {
	ConfigSource
}

func NewKafkaConfig(configSource ConfigSource) *KafkaConfig {
	return &KafkaConfig{ConfigSource: configSource}
}

func (k *KafkaConfig) update() {
	k.config()
}

// ------------------------------

// MySQLConfig MySQL 配置类
type MySQLConfig struct {
	ConfigSource
}

func NewMySQLConfig(configSource ConfigSource) *MySQLConfig {
	return &MySQLConfig{ConfigSource: configSource}
}

func (m *MySQLConfig) OutputInPlainText() string {
	panic("implement me")
}

func (m *MySQLConfig) Output() map[string]string {
	panic("implement me")
}

// ------------------------------

// ScheduledUpdater 代码热更新类
type ScheduledUpdater struct {
	initialDelayInSeconds time.Duration
	periodInSeconds       time.Duration
	Updater
}

func NewScheduledUpdater(initialDelayInSeconds time.Duration, periodInSeconds time.Duration, updater Updater) *ScheduledUpdater {
	return &ScheduledUpdater{
		initialDelayInSeconds: initialDelayInSeconds,
		periodInSeconds:       periodInSeconds,
		Updater:               updater,
	}
}

func (s *ScheduledUpdater) run() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		time.Sleep(s.initialDelayInSeconds)

		ticker := time.NewTicker(s.periodInSeconds)
		for {
			select {
			case <-ticker.C:
				s.update()
			}
		}
	}()
}

// ------------------------------

// ConfigServer 对外暴露配置信息的类
type ConfigServer struct {
	host    string
	port    int
	viewers map[string][]Viewer
	*http.Server
}

func NewConfigServer(host string, port int) *ConfigServer {
	return &ConfigServer{
		host:    host,
		port:    port,
		viewers: make(map[string][]Viewer),
		Server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
	}
}

func (c *ConfigServer) AddViewer(urlDirectory string, viewer Viewer) {
	if v, ok := c.viewers[urlDirectory]; ok {
		v = append(v, viewer)
	}

	curViewerList := make([]Viewer, 0)
	curViewerList = append(curViewerList, viewer)

	c.viewers[urlDirectory] = curViewerList

}

func (c *ConfigServer) Run() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()

		if err := c.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
