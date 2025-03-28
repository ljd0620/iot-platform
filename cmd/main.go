package main

import (
    "log"
    "net/http"

    "iot-platform/internal/mqtt"
    "iot-platform/internal/websocket"
    _ "net/http/pprof"
)

func main() {
    // Initialize MQTT client
    mqttClient := mqtt.NewMqttClient()
    err := mqttClient.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to MQTT broker: %v", err)
    }
    defer mqttClient.Disconnect()

    // Initialize WebSocket server
    wsClient := websocket.NewWebSocketClient()
    http.HandleFunc("/ws", wsClient.HandleConnections)

    go func() {
        log.Println("Starting WebSocket server on :8080")
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("Failed to start WebSocket server: %v", err)
        }
    }()

    // Start pprof server
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Keep the main goroutine running
    select {}
}