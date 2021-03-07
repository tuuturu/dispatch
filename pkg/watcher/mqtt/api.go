package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tuuturu/dispatch/pkg/core"
	"net/url"
)

func (m *mqttWatcher) RegisterHandler(handler core.Handler) {
	m.client.Subscribe(m.topic, 0, func(client mqtt.Client, message mqtt.Message) {
		err := handler.Handle(message.Payload())
		if err != nil {
			fmt.Println(err)
		}
	})
}

func (m *mqttWatcher) Open() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", m.scheme, m.hostname, m.port))
	opts.SetClientID("dispatch")

	m.client = mqtt.NewClient(opts)

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("connecting to mqtt broker: %w", token.Error())
	}

	return nil
}

func (m *mqttWatcher) Close() error {
	m.client.Disconnect(250)

	return nil
}

func NewMQTTWatcher(rawURI string) core.Watcher {
	uri, _ := url.Parse(rawURI)

	return &mqttWatcher{
		scheme:   uri.Scheme,
		hostname: uri.Hostname(),
		port:     uri.Port(),
		topic:    uri.Path,
	}
}
