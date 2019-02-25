package mqtt

import (
	"encoding/json"
	"jianghexinxi/pkg/setting"
	"jianghexinxi/pkg/util"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type DataPoint struct {
	PointId    json.Number
	Value      json.Number
	SlaveIndex json.Number
	SlaveAddr  json.Number
}

type Data struct {
	DataPoints []DataPoint
}

var con MQTT.Client

var Chan chan Data

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	var data Data
	json.Unmarshal(msg.Payload(), &data)
	Chan <- data
}

func Setup() {
	Chan = make(chan Data)
	mqttSetting := setting.MqttClientSetting
	opts := MQTT.NewClientOptions().AddBroker(mqttSetting.Host)
	opts.SetClientID(mqttSetting.ClientID)
	opts.SetUsername(mqttSetting.Username)
	password := util.EncodeMD5(mqttSetting.Password)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(f)

	con = MQTT.NewClient(opts)
	if token := con.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Mqtt Connect err %s", token.Error())
	}
}

func Subscribe() {
	topic := setting.MqttClientSetting.Topic
	if token := con.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		log.Fatal("Mqtt Subscribe err %s", token.Error())
	}

}
