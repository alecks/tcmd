// The tcmd server.
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := ":57917"
	if len(os.Args) > 1 {
		addr = ":" + os.Args[1]
	}

	l, err := net.Listen("tcp", addr)
	chk(err)
	fmt.Println("✅ listening (tcp) on", addr)
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(c)
	}
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
