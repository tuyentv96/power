package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"fmt"
	utils "apis/models/mqtt/api"
	"os"
	"time"
	conf "apis/conf"
)

var topics  = map[string]byte{

	"HTTP/#": 0,

}

var MainMqttC MainMqttClient

type MainMqttClient struct {
	client MQTT.Client
	opt *MQTT.ClientOptions
}


var mqttReceive MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())


}

func mqttOnConnect(client MQTT.Client) {
	print("connected to broker")
}

func mqttLostConnect(client MQTT.Client,err error) {
	print("disconnect to broker")


}

func MqttPublishMessage(topic string,payload []byte)  {
	MainMqttC.client.Publish(topic,0,false,payload)
}



func InitMqtt()  {
	cid:= string(string(time.Now().Unix())+utils.RandStringRunes(12))+"http"
	print("cid:",cid)
	opts := MQTT.NewClientOptions().AddBroker(conf.Mqtt_broker)
	opts.SetClientID(cid)
	opts.SetDefaultPublishHandler(mqttReceive)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(true)
	opts.SetOnConnectHandler(mqttOnConnect)
	opts.SetConnectionLostHandler(mqttLostConnect)


	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}


	if token :=c.SubscribeMultiple(topics,nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	MainMqttC.client=c
	fmt.Print("connnect ok")


}



