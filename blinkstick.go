package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

// Device type: BlinkStick
var BlinkStick DeviceType

func init() {
	BlinkStick = addDriver(usbDriver{
		Name:      "BlinkStick",
		Type:      &BlinkStick,
		VendorId:  0x20a0,
		ProductId: 0x41e5,
		Open: func(d hid.Device) (Device, error) {
			return &blinkStickDev{d}, nil
		},
	})
}

type blinkStickDev struct {
	dev hid.Device
}

func (d *blinkStickDev) SetColor(c color.Color) error {
	r, g, b, _ := c.RGBA()
	return d.dev.WriteFeature([]byte{0x01, byte(r >> 8), byte(g >> 8), byte(b >> 8)})
}
func (d *blinkStickDev) Close() {
	d.SetColor(color.Black)
	d.dev.Close()
}
