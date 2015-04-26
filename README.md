# Package to control USB-LED devices

## Supported OS
* OSX
* Windows

Linux support is planned but I didn't have a linux pc to implement the HID API.

## Supported devices

* [blink(1)](http://blink1.thingm.com/)
* [LinkM / BlinkM](http://thingm.com/products/linkm/)
* [BlinkStick](http://www.blinkstick.com/)
* [Blync](http://www.blynclight.com/)
* [Busylight UC](http://www.busylight.com/busylight-uc.html)
* [Busylight Lync](http://www.busylight.com/busylight-lync.html)
* DealExtreme USBMailNotifier
* [DreamCheeky USBMailNotifier](http://www.dreamcheeky.com/webmail-notifier)

## References

Most of the device control knowledge is taken from the [NotifierLight](http://notifierlight.blogspot.de/) project.

## Documentation
See [GoDoc](https://godoc.org/github.com/boombuler/led)

## Code example
```go
package main

import (
    "fmt"
    "github.com/boombuler/led"
    "image/color"
    "time"
)

var RED color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF}

func main() {
    for devInfo := range led.Devices() {
        dev, err := devInfo.Open()
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer func() {
            dev.SetColor(color.Black)
            dev.Close()
        }()
        dev.SetColor(RED)

        time.Sleep(2 * time.Second) // Wait 2 seconds because the device will turn off once it is closed!
    }
}
```
