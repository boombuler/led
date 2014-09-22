package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

// Device type: DreamCheeky USBMailNotifier
var DreamCheeky DeviceType

func init() {
	DreamCheeky = addDriver(usbDriver{
		Name:      "DreamCheeky USBMailNotifier",
		Type:      &DreamCheeky,
		VendorId:  0x1D34,
		ProductId: 0x0004,
		Open: func(d hid.Device) (Device, error) {
			if err := d.Write([]byte{0x00, 0x1F, 0x01, 0x29, 0x00, 0xB8, 0x54, 0x2C, 0x03}); err != nil {
				return nil, err
			}

			if err := d.Write([]byte{0x00, 0x00, 0x01, 0x29, 0x00, 0xB8, 0x54, 0x2C, 0x04}); err != nil {
				return nil, err
			}
			return &dreamCheekyDev{d}, nil
		},
	})
}

type dreamCheekyDev struct {
	dev hid.Device
}

func (d *dreamCheekyDev) SetColor(c color.Color) error {
	r, g, b, _ := c.RGBA()
	return d.dev.Write([]byte{0x00, byte(r >> 10), byte(g >> 10), byte(b >> 10), 0x00, 0x00, 0x54, 0x2C, 0x05})
}
func (d *dreamCheekyDev) Close() {
	d.SetColor(color.Black)
	d.dev.Close()
}
