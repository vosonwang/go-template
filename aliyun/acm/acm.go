package acm

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"my-project-name/util"
	"os"
)

var client config_client.IConfigClient

func Init() {
	acmConfig := constant.ClientConfig{
		AccessKey: os.Getenv("ACM_ACCESS_KEY"),
		SecretKey: os.Getenv("ACM_SECRET_KEY"),
		Endpoint:  util.GetEnv("ACM_ENDPOINT", "acm.aliyun.com:8080"),
		// 默认使用卫川阿里云ACM默认空间
		NamespaceId: util.GetEnv("ACM_NAMESPACE_ID", "0fad538c-3631-4e74-bf64-86bd1b113250"),
	}

	var err error
	client, err = clients.CreateConfigClient(map[string]interface{}{
		constant.KEY_CLIENT_CONFIG: acmConfig,
	})
	if err != nil {
		log.Fatalf("err: %v,config: %+v", err, acmConfig)
	}
}

func GetConfig(dataId, group string) (string, error) {
	// 从阿里云获取应用配置
	return client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})
}

func ListenConfig(dataId, group string, listener vo.Listener) error {
	return client.ListenConfig(vo.ConfigParam{
		DataId:   dataId,
		Group:    group,
		OnChange: listener,
	})
}
