// The tcmd client.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"tcmd/internal/definitions"
)

func main() {
	port := "57917"
	if len(os.Args) < 2 {
		fmt.Println("❌ required: host")
		return
	} else if len(os.Args) > 2 {
		port = os.Args[2]
	}

	c, err := net.Dial("tcp", os.Args[1]+":"+port)
	chk(err)
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("# ")
		req, _ := reader.ReadString('\n')
		marshalled, _ := json.Marshal(definitions.TcmdRequest{
			Method: strings.TrimSpace(req),
		})
		c.Write(append(marshalled, '\n'))

		msg, err := bufio.NewReader(c).ReadString('\n')
		chk(err)
		fmt.Print(msg)
	}
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
