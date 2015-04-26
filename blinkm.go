package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

// Device type: LinkM / BlinkM
var BlinkM DeviceType

func init() {
	BlinkM = addDriver(usbDriver{
		Name:      "LinkM / BlinkM",
		Type:      &BlinkM,
		VendorId:  0x20A0,
		ProductId: 0x4110,
		Open: func(d hid.Device) (Device, error) {
			return &simpleHidDevice{
				device:     d,
				setColorFn: blinkMSetColor,
			}, nil
		},
	})
}

func blinkMSetColor(d hid.Device, c color.Color) error {
	r, g, b, _ := c.RGBA()
	return d.WriteFeature([]byte{0x01, 0xDA, 0x01, 0x05, 0x00, 0x09, 0x6E, byte(r >> 8), byte(g >> 8), byte(b >> 8)})
}
