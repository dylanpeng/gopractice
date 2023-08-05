package main

import "sync"

var (
	Cfg *Configuration
	mu  sync.RWMutex
)

type AppConfig struct {
	Secret           string                `json:"secret" default:""`
	ChainId          int                   `json:"chain" default:"56" mapstructure:"chain"`
	Model            string                `json:"model" default:"main_net"`
	Env              string                `json:"env" default:""`
	ConfigUrl        string                `json:"config_url" default:"" mapstructure:"config_url"`
	Cros             bool                  `json:"cros" default:"false"`
	Strings          []string              `json:"strings"`
	Ints             []int                 `json:"ints"`
	ServerConfig     *ServerConfig         `json:"server_config" mapstructure:"server_config"` // server_config.host = host
	MapStr           map[string]string     `json:"map_str" mapstructure:"map_str"`             // map_str.a = a
	ServerConfigList []*ServerConfig       `json:"server_list" mapstructure:"server_list"`     //无法实现
	ServerConfigMap  map[int]*ServerConfig `json:"server_map" mapstructure:"server_map"`       //server_map.11.port
	//SyncOrder bool   `json:"sync_order" default:"false"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type MysqlConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type Configuration struct {
	App   *AppConfig   `json:"app"` //
	Mysql *MysqlConfig `json:"mysql"`
}
