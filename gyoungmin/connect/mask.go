package connect

import (
	"fmt"
	"net"
	"os"
)

func mask() {
	var IP string

	fmt.Scanf("%s", IP)

	if len(IP) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", IP)
		os.Exit(1)
	}

	dotAddr := IP

	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

}
