package main

import (
	"fmt"

	"github.com/ximoqing/goamap"
)

func main() {
	r := goamap.IpJson("9ae6604ba94f0296683e556e74f3f63d", "114.247.50.2")
	fmt.Println(r)
}
