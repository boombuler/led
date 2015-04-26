package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

// Device type: DealExtreme USBMailNotifier
var DealExtreme DeviceType

func init() {
	DealExtreme = addDriver(usbDriver{
		Name:      "DealExtreme USBMailNotifier",
		Type:      &DealExtreme,
		VendorId:  0x1294,
		ProductId: 0x1320,
		Open: func(d hid.Device) (Device, error) {
			return &simpleHidDevice{
				device:     d,
				setColorFn: dealExtremeSetColor,
			}, nil
		},
	})
}

func dealExtremeSetColor(d hid.Device, c color.Color) error {
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
	return d.Write([]byte{0x00, byte(palette.Index(c))})
}
