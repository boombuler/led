package led

import (
	"github.com/boombuler/hid"
)

type usbDriver struct {
	Name      string
	Type      *DeviceType
	VendorId  uint16
	ProductId uint16
	Open      func(hidDevice hid.Device) (Device, error)
}

func (drv usbDriver) name() string {
	return drv.Name
}

func (sd usbDriver) convert(hDev *hid.DeviceInfo) DeviceInfo {
	if hDev.VendorId == sd.VendorId && hDev.ProductId == sd.ProductId {
		return &usbDeviceInfo{hDev, &sd}
	}
	return nil
}

type usbDeviceInfo struct {
	*hid.DeviceInfo
	driver *usbDriver
}

func (usb usbDeviceInfo) GetType() DeviceType {
	return *usb.driver.Type
}

func (usb usbDeviceInfo) GetPath() string {
	return usb.Path
}

func (usb *usbDeviceInfo) Open() (Device, error) {
	d, err := usb.DeviceInfo.Open()
	if err != nil {
		return nil, err
	} else {
		return usb.driver.Open(d)
	}
}
