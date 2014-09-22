package led

import (
	"fmt"
	"github.com/boombuler/hid"
	"image/color"
)

type DeviceType int

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

type Device interface {
	SetColor(c color.Color) error
	Close()
}

type DeviceInfo interface {
	GetPath() string
	GetType() DeviceType
	Open() (Device, error)
}

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

func ByPath(path string) (DeviceInfo, error) {
	hd := hid.ByPath(path)
	if hd != nil {
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
