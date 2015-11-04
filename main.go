package main

import (
	"fmt"

	"github.com/GeertJohan/go.hid"
	"github.com/HackerLoop/rotonde-client.go"
	"github.com/HackerLoop/rotonde/shared"
)

const charMap string = "    abcdefghijklmnopqrstuvwxyz1234567890\n  \t -=[]\\ ;'`,./                           /*-+\n123467890.\\="
const charMapMaj string = "    ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()\n  \t _+{}| :\"~<>?                           /*-+\n123467890.\\="

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

func startListening(c chan string, device *hid.Device) {
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
		// PrintHex(b, n)
		if b[2] != 0 {
			if b[2] == 0x51 {
				fmt.Println(current)
				c <- current
				current = ""
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

	c := make(chan string, 10)
	go startListening(c, device)

	r := client.NewClient("ws://127.0.0.1:4224/")

	event := &rotonde.Definition{"BARCODE_RECEIVED", "event", rotonde.FieldDefinitions{}}
	event.PushField("code", "string", "")
	r.AddLocalDefinition(event)

	go func() {
		for code := range c {
			r.SendEvent("BARCODE_RECEIVED", rotonde.Object{
				"code": code,
			})
		}
	}()

	select {}
}
