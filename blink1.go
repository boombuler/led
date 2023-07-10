package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

// Device type: blink(1)
var Blink1 DeviceType

func init() {
	Blink1 = addDriver(usbDriver{
		Name:      "blink(1)",
		Type:      &Blink1,
		VendorId:  0x27B8,
		ProductId: 0x01ED,
		Open: func(d hid.Device) (Device, error) {
			return &simpleHidDevice{
				device:     d,
				setColorFn: blink1SetColor,
			}, nil
		},
	})
}

func blink1SetColor(d hid.Device, c color.Color) error {
	r, g, b, _ := c.RGBA()
	return d.WriteFeature([]byte{0x01, 0x63, byte(r >> 8), byte(g >> 8), byte(b >> 8), 0x00, 0x00, 0x00, 0x00})
}
