package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

var DealExtreme DeviceType

func init() {
	DealExtreme = addDriver(usbDriver{
		Name:      "DealExtreme USBMailNotifier",
		Type:      &DealExtreme,
		VendorId:  0x1294,
		ProductId: 0x1320,
		Open: func(d hid.Device) (Device, error) {
			return &dealExtremeDev{d}, nil
		},
	})
}

type dealExtremeDev struct {
	dev hid.Device
}

func (d *dealExtremeDev) SetColor(c color.Color) error {
	palette := color.Palette{
		color.RGBA{0x00, 0x00, 0x00, 0x00},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
	}
	return d.dev.Write([]byte{0x00, byte(palette.Index(c))})
}
func (d *dealExtremeDev) Close() {
	d.SetColor(color.Black)
	d.dev.Close()
}
