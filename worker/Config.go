package worker

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort         int      `json:"apiPort"`
	ApiReadTimeout  int      `json:"apiReadTimeout"`
	ApiWriteTimeout int      `json:"apiWriteTimeout"`
	EtcdEndpoints   []string `json:"etcdEndpoints"`
	EtcdDialTimeout int      `json:"etcdDialTimeout"`
	WebRoot         string   `json:"webroot"`
}

var (
	// 单例对象
	G_config *Config
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)

	// 1、读取配置文件
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	// 2、JSON反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3、赋值单例
	G_config = &conf

	return
}
