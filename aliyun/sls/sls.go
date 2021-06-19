package sls

import (
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"time"
)

type Config struct {
	SLSAccessKeyID     string
	SLSAccessKeySecret string
	SLSEndpoint        string
	SLSProject         string
	SLSLogStore        string
}

var (
	producerInstance *producer.Producer
	// logStore可以通过Config设置，也可以通过SetLogStore
	project, logStore string
)

func SetLogStore(store string) {
	logStore = store
}

func Init(parseConfig func(v interface{}) error) error {
	var cfg Config

	if err := parseConfig(&cfg); err != nil {
		return err
	}

	project = cfg.SLSProject

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = cfg.SLSEndpoint
	producerConfig.AccessKeyID = cfg.SLSAccessKeyID
	producerConfig.AccessKeySecret = cfg.SLSAccessKeySecret

	producerInstance = producer.InitProducer(producerConfig)

	producerInstance.Start()

	return nil
}

func SendLog(topic, source string, t time.Time, addLogMap map[string]string) error {
	log := producer.GenerateLog(uint32(t.Unix()), addLogMap)
	return producerInstance.SendLog(project, logStore, topic, source, log)
}
