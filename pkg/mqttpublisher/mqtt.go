package mqttpublisher

import (
	mqttPkg "github.com/eclipse/paho.mqtt.golang"
	"log"
	"yumikokawaii.iot.com/config"
)

type MQTTClient struct {
	client mqttPkg.Client
	qos    byte
}

func NewMQTTClient(config *config.AppConfig) *MQTTClient {

	mqttCfg := config.MqttCfg
	clientOpts := mqttPkg.NewClientOptions()
	clientOpts.AddBroker(mqttCfg.Broker)
	clientOpts.SetClientID(mqttCfg.ClientID)
	client := mqttPkg.NewClient(clientOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Println("[MQTT Client] - Connected!")
	return &MQTTClient{
		client: client,
		qos:    mqttCfg.QOS,
	}
}

func (m *MQTTClient) Send(topic string, data Serializable) error {
	if token := m.client.Publish(topic, m.qos, false, data.ToBytes()); token.Wait() && token.Error() != nil {
		log.Printf("[MQTT Client] - Error sending message to MQTT server: %v", token.Error())
		return token.Error()
	}
	log.Printf("[MQTT Client] - Send message success, content: %v", data)
	return nil
}
