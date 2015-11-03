package main

import (
	"fmt"

	"github.com/GeertJohan/go.hid"
)

const charMap string = "    abcdefghijklmnopqrstuvwxyz1234567890\n  \t -=[]\\ ;'`,./"
const charMapMaj string = "    ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()\n  \t _+{}| :\"~<>?"

func PrintHex(buffer []byte, n int) {
	l := ""
	for i, b := range buffer {
		if i > 0 {
			l += ":"
		}
		l += fmt.Sprintf("%.02x", b)
	}
	fmt.Println(l)
}

func startListening(device *hid.Device) {
	b := make([]byte, 20)
	current := ""

	fmt.Println("Start listening HID")
	for {
		n, err := device.ReadTimeout(b, 20)
		if err != nil {
			panic(err)
		}
		if n == 0 {
			continue
		}
		PrintHex(b, n)
		if b[2] != 0 {
			if b[2] == 0x51 {
				fmt.Println(current)
				continue
			} else if b[2] == 0x28 {
				continue
			}
			index := int(b[2])
			var cm string
			if b[0] == 0x02 {
				cm = charMapMaj
			} else {
				cm = charMap
			}
			if index > 0 && index < len(cm) {
				current += fmt.Sprintf("%s", string(cm[index]))
			} else {
				fmt.Printf("unknown key %d 0x%x\n", index, index)
			}
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
