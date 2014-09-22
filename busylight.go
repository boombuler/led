package led

import (
	"github.com/boombuler/hid"
	"image/color"
	"time"
)

// Device type: BusyLight UC
var BusyLightUC DeviceType

// Device type: BusyLight Lync
var BusyLightLync DeviceType

func init() {
	BusyLightUC = addDriver(usbDriver{
		Name:      "BusyLight UC",
		Type:      &BusyLightUC,
		VendorId:  0x27BB,
		ProductId: 0x3BCB,
		Open: func(d hid.Device) (Device, error) {
			return newBusyLight(d, func(c color.Color) {
				r, g, b, _ := c.RGBA()
				d.Write([]byte{0x00, 0x00, 0x00, byte(r >> 8), byte(g >> 8), byte(b >> 8), 0x00, 0x00, 0x00})
			}), nil
		},
	})
	BusyLightLync = addDriver(usbDriver{
		Name:      "BusyLight Lync",
		Type:      &BusyLightLync,
		VendorId:  0x04D8,
		ProductId: 0xF848,
		Open: func(d hid.Device) (Device, error) {
			return newBusyLight(d, func(c color.Color) {
				r, g, b, _ := c.RGBA()
				d.Write([]byte{0x8A, 0x00, 0x00, byte(r >> 8), byte(g >> 8), byte(b >> 8), 0x01, 0x00, 0x80})
			}), nil
		},
	})
}

type busylightDev struct {
	closeChan chan<- struct{}
	colorChan chan<- color.Color
}

func newBusyLight(d hid.Device, setcolorFn func(c color.Color)) *busylightDev {
	closeChan := make(chan struct{})
	colorChan := make(chan color.Color)
	ticker := time.NewTicker(20 * time.Second) // If nothing is send after 30 seconds the device turns off.
	go func() {
		var curColor color.Color = color.Black
		closed := false
		for !closed {
			select {
			case <-ticker.C:
				setcolorFn(curColor)
			case col := <-colorChan:
				curColor = col
				setcolorFn(curColor)
			case <-closeChan:
				ticker.Stop()
				setcolorFn(color.Black) // turn off device
				d.Close()
				closed = true
			}
		}
	}()
	return &busylightDev{closeChan: closeChan, colorChan: colorChan}

}

func (d *busylightDev) SetColor(c color.Color) error {
	d.colorChan <- c
	return nil
}
func (d *busylightDev) Close() {
	d.closeChan <- struct{}{}
}
