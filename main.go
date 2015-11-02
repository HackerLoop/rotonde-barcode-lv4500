package main

import(
	"fmt"

	"github.com/GeertJohan/go.hid"
)

func main() {
	list, err := hid.Enumerate(0x0, 0x0)
	if err != nil {
		panic(err)
	}

	for _, item := range list {
		if item.VendorId == 0x1eab && item.ProductId == 0x8203 {
			fmt.Printf("%s %s\n", item.Manufacturer, item.Product)
			_, err := item.Device()
			if err != nil {
				panic(err)
			}
			break
		}
	}

}
