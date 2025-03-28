package device

import (
	"sync"
)

// Device represents a connected IoT device.
type Device struct {
	ID   string
	Name string
}

// DeviceManager manages the connected devices.
type DeviceManager struct {
	devices map[string]Device
	mu      sync.Mutex
}

// NewDeviceManager creates a new DeviceManager.
func NewDeviceManager() *DeviceManager {
	return &DeviceManager{
		devices: make(map[string]Device),
	}
}

// AddDevice adds a new device to the manager.
func (dm *DeviceManager) AddDevice(device Device) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	dm.devices[device.ID] = device
}

// RemoveDevice removes a device from the manager by ID.
func (dm *DeviceManager) RemoveDevice(deviceID string) {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	delete(dm.devices, deviceID)
}

// ListDevices returns a slice of all connected devices.
func (dm *DeviceManager) ListDevices() []Device {
	dm.mu.Lock()
	defer dm.mu.Unlock()
	devicesList := make([]Device, 0, len(dm.devices))
	for _, device := range dm.devices {
		devicesList = append(devicesList, device)
	}
	return devicesList
}