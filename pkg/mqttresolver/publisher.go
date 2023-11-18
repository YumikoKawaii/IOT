package mqttresolver

import (
	mqttPkg "github.com/eclipse/paho.mqtt.golang"
	"log"
	"yumikokawaii.iot.com/config"
)

type Publisher struct {
	client mqttPkg.Client
	qos    byte
}

func NewPublisher(config *config.AppConfig) *Publisher {

	mqttCfg := config.MqttPublisherCfg
	clientOpts := mqttPkg.NewClientOptions()
	clientOpts.AddBroker(mqttCfg.PBroker)
	clientOpts.SetClientID(mqttCfg.PClientID)
	client := mqttPkg.NewClient(clientOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Println("[MQTT Publisher] - Connected!")
	return &Publisher{
		client: client,
		qos:    mqttCfg.PQOS,
	}
}

func (m *Publisher) Send(topic string, data Serializable) error {
	if token := m.client.Publish(topic, m.qos, false, data.ToBytes()); token.Wait() && token.Error() != nil {
		log.Printf("[MQTT Client] - Error sending message to MQTT server: %v", token.Error())
		return token.Error()
	}
	log.Printf("[MQTT Client] - Send message success, content: %v", data)
	return nil
}
