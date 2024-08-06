package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"xiongmao/config"
)

var ClientVar MQTT.Client

func Init() {

	mqttConfig := config.MqttVar

	opts := MQTT.NewClientOptions().AddBroker(mqttConfig.Server)

	opts.SetClientID(mqttConfig.ClientID)
	opts.SetUsername(mqttConfig.Username)
	opts.SetPassword(mqttConfig.Password)

	// 设置客户端连接时的回调
	opts.OnConnect = func(client MQTT.Client) {
		fmt.Println("MQTT Connected")
	}

	// 设置收到消息时的回调
	opts.OnConnectionLost = func(client MQTT.Client, err error) {
		fmt.Printf("MQTT Connection lost：%v\n", err)
	}

	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("MQTT Received message on topic：%s\nMessage: %s\n", msg.Topic(), msg.Payload())
	}

	ClientVar = MQTT.NewClient(opts)

	if token := ClientVar.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("MQTT Error connecting to broker: %v\n", token.Error())
	}

	topics := strings.Split(mqttConfig.Topic, ",")

	for _, topic := range topics {
		if token := ClientVar.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
			fmt.Printf("MQTT Error subscribing to topic %s: %v\n", topic, token.Error())
		} else {
			fmt.Printf("MQTT Subscribed to topic: %s\n", topic)
		}
	}
}

// Publish 发送消息
func Publish(topic string, message string) {

	token := ClientVar.Publish(topic, 0, false, message)

	token.Wait()

	if token.Error != nil {

		fmt.Printf("Error publishing to topic %s: %v\n", topic, token.Error())

		return
	}

	fmt.Printf("Published message to topic %s: %s\n", topic, message)
}

// PublishQos 发送消息
func PublishQos(topic string, message string, qos byte) {

	token := ClientVar.Publish(topic, qos, false, message)

	token.Wait()

	if token.Error != nil {

		fmt.Printf("Error publishing to topic %s: %v\n", topic, token.Error())

		return
	}

	fmt.Printf("Published message to topic %s: %s\n", topic, message)
}
