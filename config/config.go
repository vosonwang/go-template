package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"my-project-name/aliyun/acm"
	"my-project-name/util"
)

type InitFun func(parseConfig func(v interface{}) error) error

var (
	host, env, projectName, version string
	port                            int
)

// executeInitFun 执行各项初始化方法
func executeInitFun(data string, fs []InitFun) error {
	for _, f := range fs {
		if err := f(func(v interface{}) error {
			_, err := toml.Decode(data, v)
			return err
		}); err != nil {
			return err
		}
	}
	return nil
}

func Init(project string, fs ...InitFun) {
	projectName = project

	group := fmt.Sprintf("%v_%v", util.GetEnv("GROUP", "RICNSMART"), Env())

	// 从ACM中获取项目配置
	content, err := acm.GetConfig(projectName, group)
	if err != nil {
		log.Fatal(err)
	}

	if err := executeInitFun(content, fs); err != nil {
		log.Fatal(err)
	}

	// 监听配置变更
	go func() {
		err := acm.ListenConfig(projectName, group, func(namespace, group, dataId, data string) {
			if err := executeInitFun(data, fs); err != nil {
				log.Printf("配置更新失败: %v", err)
			}
		})

		if err != nil {
			log.Fatal(err)
		}
	}()
}

func Host() string {
	if host == "" {
		var err error
		host, err = util.GetPublicIP()
		if err != nil {
			log.Fatal(err)
		}
	}

	return host
}

func HttpPort() int {
	if port == 0 {
		port = util.GetInt("HTTP_PORT", 1234)
	}
	return port
}

func Env() string {
	if env == "" {
		env = util.GetEnv("ENV", "DEVELOP")
	}
	return env
}

func IsDev() bool {
	if Env() == "DEVELOP" {
		return true
	}
	return false
}

func ProjectName() string {
	return projectName
}

func Version() string {
	if version == "" {
		version = util.GetEnv("VERSION", "v0.0.0")
	}
	return version
}

func Debug() bool {
	return util.GetBool("DEBUG")
}
