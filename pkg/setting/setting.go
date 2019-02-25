package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type MqttClient struct {
	Host     string
	ClientID string
	Username string
	Password string
	Topic    string
}

var MqttClientSetting = &MqttClient{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("E:/go_code/src/jianghexinxi/conf/app.ini")
	if err != nil {
		log.Fatal("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	if err != nil {
		log.Fatal("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("mqtt", MqttClientSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal("Cfg.MapTo RedisSetting err: %v", err)
	}
}
