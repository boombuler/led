package led

import (
	"github.com/boombuler/hid"
	"image/color"
)

// Device type: Blync
var Blync DeviceType

func init() {
	Blync = addDriver(blyncDriver{
		usbDriver{
			Name:      "Blync",
			Type:      &Blync,
			VendorId:  0x1130,
			ProductId: 0x0001,
			Open: func(d hid.Device) (Device, error) {
				return &simpleHidDevice{
					device:     d,
					setColorFn: blyncDevSetColor,
				}, nil
			},
		},
	})
}

type blyncDriver struct {
	usbDriver
}

func (drv blyncDriver) convert(hDev *hid.DeviceInfo) DeviceInfo {
	// blync adds two devices. but only the one which accepts feature reports will work.
	if hDev.FeatureReportLength == 0 {
		return drv.usbDriver.convert(hDev)
	}
	return nil
}

func blyncDevSetColor(d hid.Device, c color.Color) error {
	palette := color.Palette{
		color.RGBA{0x00, 0x00, 0x00, 0x00}, // black
		color.RGBA{0xff, 0xff, 0xff, 0xff}, // white
		color.RGBA{0x00, 0xff, 0xff, 0xff}, // cyan
		color.RGBA{0xff, 0x00, 0xff, 0xff}, // magenta
		color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
		color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
		color.RGBA{0x00, 0xff, 0x00, 0xff}, // lime
		color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
	}

	value := byte((palette.Index(c) * 16) + 127)
	return d.Write([]byte{0x00, 0x55, 0x53, 0x42, 0x43, 0x00, 0x40, 0x02, value})
}
