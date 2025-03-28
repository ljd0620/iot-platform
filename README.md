# IoT Platform

This project is an IoT platform designed to support massive device connectivity using WebSocket and MQTT protocols for communication with clients.

## Features

- **Device Connectivity**: Supports a large number of devices connecting simultaneously.
- **WebSocket Communication**: Enables real-time communication with clients.
- **MQTT Protocol**: Efficiently manages device messaging and communication.

## Project Structure

```
iot-platform
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── mqtt
│   │   └── client.go    # MQTT client management
│   ├── websocket
│   │   └── client.go    # WebSocket client management
│   └── device
│       └── manager.go   # Device management
├── pkg
│   └── config
│       └── config.go    # Configuration management
├── go.mod                # Module definition
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the Repository**:
   ```
   git clone <repository-url>
   cd iot-platform
   ```

2. **Install Dependencies**:
   ```
   go mod tidy
   ```

3. **Run the Application**:
   ```
   go run cmd/main.go
   ```

## Usage Guidelines

- Ensure that your devices are configured to connect using either MQTT or WebSocket protocols.
- Use the provided methods in the `DeviceManager` to manage device states.
- Configure the application settings in the `config.go` file as needed.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.