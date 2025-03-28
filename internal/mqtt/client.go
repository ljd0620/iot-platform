package mqtt

import (
	"fmt"
	"log"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttClient struct {
	client mqtt.Client
	mu     sync.Mutex
}

func NewMqttClient(broker string, clientID string) *MqttClient {
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID)
	client := mqtt.NewClient(opts)
	return &MqttClient{client: client}
}

func (m *MqttClient) Connect() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (m *MqttClient) Publish(topic string, payload interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	token := m.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func (m *MqttClient) Subscribe(topic string, callback mqtt.MessageHandler) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	token := m.client.Subscribe(topic, 0, callback)
	token.Wait()
	return token.Error()
}