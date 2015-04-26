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
			return &simpleHidDevice{
				device:     d,
				setColorFn: blinkStickSetColor,
			}, nil
		},
	})
}

func blinkStickSetColor(d hid.Device, c color.Color) error {
	r, g, b, _ := c.RGBA()
	return d.WriteFeature([]byte{0x01, byte(r >> 8), byte(g >> 8), byte(b >> 8)})
}
