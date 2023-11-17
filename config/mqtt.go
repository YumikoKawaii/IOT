package config

type MQTTConfig struct {
	Broker   string `default:"tcp://test.mosquitto.org:1883" env:"MQTT_BROKER"`
	ClientID string `default:"yumiko.publisher.sv" env:"CLIENT_ID"`
	QOS      byte   `default:"2" env:"QOS"`
}
