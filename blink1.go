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
			return &blink1Dev{d}, nil
		},
	})
}

type blink1Dev struct {
	dev hid.Device
}

func (d *blink1Dev) SetColor(c color.Color) error {
	r, g, b, _ := c.RGBA()
	return d.dev.WriteFeature([]byte{0x01, 0x63, byte(r >> 8), byte(g >> 8), byte(b >> 8), 0x00, 0x00, 0x00, 0x00})
}
func (d *blink1Dev) Close() {
	d.dev.Close()
}
