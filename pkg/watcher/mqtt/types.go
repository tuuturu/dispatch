package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

type mqttWatcher struct {
	client mqtt.Client

	scheme   string
	hostname string
	port     string

	topic string
}
