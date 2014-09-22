package led

import (
	"fmt"
	"github.com/boombuler/hid"
	"image/color"
)

// Device type identifies the device type. the IDs may change on each program start.
type DeviceType int

// String returns a representation of the device type
func (dt DeviceType) String() string {
	idx := int(dt)
	if idx < 0 || idx >= len(drivers) {
		return ""
	}
	return drivers[idx].name()
}

type driver interface {
	name() string
	convert(hDev *hid.DeviceInfo) DeviceInfo
}

var drivers = []driver{}

func addDriver(drv driver) DeviceType {
	dt := DeviceType(len(drivers))
	drivers = append(drivers, drv)
	return dt
}

// Device is an opened LED device.
type Device interface {
	// SetColor sets the color of the LED to the closest supported color.
	SetColor(c color.Color) error
	// Close the device and release all resources
	Close()
}

// DeviceInfo keeps information about a physical LED device
type DeviceInfo interface {
	// GetPath returns a system specific path which can be used to find the device
	GetPath() string
	// GetType returns the "driver type" of the device
	GetType() DeviceType
	// Open opens the device for usage
	Open() (Device, error)
}

// Devices returns a channel with all connected LED devices
func Devices() <-chan DeviceInfo {
	result := make(chan DeviceInfo)
	go func() {
		for d := range hid.Devices() {
			di := toLedDeviceInfo(d)
			if di != nil {
				result <- di
			}
		}
		close(result)
	}()
	return result
}

// ByPath searches a device by given system specific path. (The path can be obtained from the GetPath func from DeviceInfo)
func ByPath(path string) (DeviceInfo, error) {
	hd, err := hid.ByPath(path)
	if err != nil {
		return nil, err
	}
	led := toLedDeviceInfo(hd)
	if led == nil {
		return nil, fmt.Errorf("Unknown LED device (VID: %v, PID: %v)", hd.VendorId, hd.ProductId)
	}
	return led, nil
}

func toLedDeviceInfo(dev *hid.DeviceInfo) DeviceInfo {
	for _, drv := range drivers {
		di := drv.convert(dev)
		if di != nil {
			return di
		}
	}
	return nil
}
