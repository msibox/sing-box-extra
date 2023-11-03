package main

import (
	"fmt"
	_ "unsafe"

	"github.com/msibox/sing-box-extra/boxbox"
	"github.com/msibox/sing-box-extra/boxmain"
	_ "github.com/msibox/sing-box-extra/distro/all"
)

func main() {
	fmt.Println("sing-box-extra:", boxbox.Version)
	fmt.Println()

	// sing-box
	boxmain.Main()
}
