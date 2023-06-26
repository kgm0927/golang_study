package connect

import (
	"fmt"
	"net"
	"os"
)

func IP_make() {

	var IP string
	fmt.Scanf("%s", &IP)

	if len(IP) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", IP)
		os.Exit(1)
	}
	name := IP

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invaild address")

	} else {
		fmt.Println("The address is", addr.String())
	}
	os.Exit(0)
}
