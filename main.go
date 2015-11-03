package main

import (
	"fmt"

	"github.com/GeertJohan/go.hid"
)

const charMap string = "    abcdefghijklmnopqrstuvwxyz1234567890\n  \t -=[]\\ ;'`,./"
const charMapMaj string = "    ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()\n  \t _+{}| :\"~<>?"

func startListening(device *hid.Device) {
	b := make([]byte, 20)

	fmt.Println("Start listening HID")
	for {
		n, err := device.ReadTimeout(b, 20)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			continue
		}
		if b[2] != 0 {
			index := int(b[2])
			if index > 0 && index < len(charMap) {
				fmt.Printf("%s", string(charMap[index]))
			} else {
				fmt.Printf("unknown key %d 0x%x\n", index, index)
			}
		} else {
			for _, b := range b[:n] {
				fmt.Printf("%x", b)
			}
			fmt.Println()
		}
	}

}

func main() {
	list, err := hid.Enumerate(0x0, 0x0)
	if err != nil {
		panic(err)
	}

	var device *hid.Device
	for _, item := range list {
		if item.VendorId == 0x1eab && item.ProductId == 0x8203 {
			fmt.Printf("%s %s\n", item.Manufacturer, item.Product)
			device, err = item.Device()
			if err != nil {
				panic(err)
			}
			break
		}
	}
	if device == nil {
		return
	}

	go startListening(device)
	select {}

}
