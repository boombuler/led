package led

import (
	"errors"
	"github.com/boombuler/hid"
	"image/color"
)

type simpleHidDevice struct {
	device     hid.Device
	keepAlive  bool
	setColorFn func(d hid.Device, c color.Color) error
}

// Error: Device was already closed and is not ready for interaction anymore
var ErrDeviceClosed = errors.New("Device is already closed")

// SetKeepActive sets a value that tells the device not turn off the device on calling Close
func (s *simpleHidDevice) SetKeepActive(v bool) error {
	if s.device != nil {
		s.keepAlive = v
		return nil
	}
	return ErrDeviceClosed
}

// Close the device and release all resources
func (s *simpleHidDevice) Close() {
	if s.device != nil {
		if !s.keepAlive {
			s.SetColor(color.Black)
		}
		s.device.Close()
		s.device = nil
	}

}

// SetColor sets the color of the LED to the closest supported color.
func (s *simpleHidDevice) SetColor(c color.Color) error {
	if s.device != nil {
		return s.setColorFn(s.device, c)
	}
	return ErrDeviceClosed
}
