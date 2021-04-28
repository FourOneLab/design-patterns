package design_principles

import (
	"testing"
	"time"
)

func TestApplication(t *testing.T) {
	configSource := NewZookeeperConfigSource()

	redisCfg := NewRedisConfig(configSource)
	redisCfg.config()

	kafkaCfg := NewKafkaConfig(configSource)
	kafkaCfg.config()

	mysqlCfg := NewMySQLConfig(configSource)
	mysqlCfg.config()

	redisScheduledUpdater := NewScheduledUpdater(10*time.Second, 60*time.Second, redisCfg)
	redisScheduledUpdater.run()

	kafkaScheduleUpdater := NewScheduledUpdater(10*time.Second, 60*time.Second, kafkaCfg)
	kafkaScheduleUpdater.run()

	configServer := NewConfigServer("127.0.0.1", 2389)
	configServer.AddViewer("/config", redisCfg)
	configServer.AddViewer("/config", mysqlCfg)
	configServer.Run()
}
