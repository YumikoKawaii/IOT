package config

type MQTTPublisherConfig struct {
	PBroker   string `default:"tcp://test.mosquitto.org:1883" env:"MQTT_BROKER"`
	PClientID string `default:"yumiko.publisher.sv" env:"CLIENT_ID"`
	PQOS      byte   `default:"2" env:"QOS"`
}

type MQTTSubscriberConfig struct {
	SBroker   string `default:"tcp://test.mosquitto.org:1883" env:"MQTT_BROKER"`
	SClientID string `default:"yumiko.subscriber.client" env:"CLIENT_ID"`
	StatTopic string `default:"yumiko.iot.device.stat" env:"DEVICE_STAT_TOPIC"`
	SQOS      byte   `default:"2" env:"QOS"`
}
