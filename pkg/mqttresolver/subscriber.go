package mqttresolver

import (
	mqttPkg "github.com/eclipse/paho.mqtt.golang"
	"log"
	"yumikokawaii.iot.com/config"
)

type Subscriber struct {
	client mqttPkg.Client
	topic  string
	qos    byte
}

type MessageHandlerFn func(message mqttPkg.Message) error

func NewSubscriber(config *config.AppConfig) *Subscriber {

	cfg := config.MqttSubscriberCfg

	opts := mqttPkg.NewClientOptions()
	opts.AddBroker(cfg.SBroker)
	opts.SetClientID(cfg.SClientID)

	client := mqttPkg.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	log.Println("[MQTT Stat Subscriber] - Connected!")
	return &Subscriber{
		client: client,
		topic:  cfg.StatTopic,
		qos:    cfg.SQOS,
	}
}

func (s *Subscriber) Consume(fn MessageHandlerFn) {
	if !s.client.IsConnected() {
		log.Println("[MQTT Subscriber] - Disconnected!")
		return
	}

	s.client.Subscribe(s.topic, s.qos, func(client mqttPkg.Client, message mqttPkg.Message) {
		if err := fn(message); err != nil {
			log.Printf("[MQTT Subscriber] - Error handle message: %v\n", err.Error())
		}
	})

}
