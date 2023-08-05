package main

import (
	"bytes"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/spf13/viper"
	"time"
)

var nameSpaceMap = make(map[string]agollo.Client)
var appConf = &AppConfig{}

func main() {
	c := &config.AppConfig{
		AppID:          "aaaaaaaa",
		Cluster:        "default",
		IP:             "http://localhost:8080",
		NamespaceName:  "application",
		IsBackupConfig: false,
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功")
	nameSpaceMap["application"] = client

	content := client.GetConfig(c.NamespaceName).GetContent()
	conf := viper.New()
	conf.SetConfigType("properties")
	err := conf.ReadConfig(bytes.NewBufferString(content))
	if err != nil {
		fmt.Println(err)
	}

	err = conf.Unmarshal(appConf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", *appConf)

	client.AddChangeListener(&Listerner{})

	fmt.Println(content)
	//Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	value, _ := cache.Get("ccccc")
	fmt.Println(value)

	time.Sleep(10 * time.Minute)
	return
}

type Listerner struct {
}

// OnChange 增加变更监控
func (l *Listerner) OnChange(event *storage.ChangeEvent) {
	if v, ok := nameSpaceMap[event.Namespace]; ok {
		content := v.GetConfig(event.Namespace).GetContent()
		conf := viper.New()
		conf.SetConfigType("properties")
		err := conf.ReadConfig(bytes.NewBufferString(content))
		if err != nil {
			fmt.Println(err)
		}

		err = conf.Unmarshal(appConf)
		if err != nil {
			fmt.Println(err)
		}
	}

}

// OnNewestChange 监控最新变更
func (l *Listerner) OnNewestChange(event *storage.FullChangeEvent) {
	if v, ok := nameSpaceMap[event.Namespace]; ok {
		content := v.GetConfig(event.Namespace).GetContent()
		conf := viper.New()
		conf.SetConfigType("properties")
		err := conf.ReadConfig(bytes.NewBufferString(content))
		if err != nil {
			fmt.Println(err)
		}

		err = conf.Unmarshal(appConf)
		if err != nil {
			fmt.Println(err)
		}
	}
}
