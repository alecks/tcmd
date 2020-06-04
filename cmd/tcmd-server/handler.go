package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"tcmd/internal/definitions"
)

func handle(c net.Conn) {
	addr := c.RemoteAddr().String()
	fmt.Printf("👋 serving %s\n", addr)

	for {
		data, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			txt := err.Error()
			if txt == "EOF" {
				fmt.Println("❌ forcefully closed", addr)
				return
			}

			fmt.Println(txt)
			return
		}

		var res definitions.TcmdRequest
		if err := json.Unmarshal(data, &res); err != nil {
			c.Write(append(definitions.JsonErrMarshalled, '\n'))
			continue
		}
		if res.Method == definitions.CloseMethod {
			break
		}

		// TODO(fjah): Allow this to be customised so that runtimes etc. can be used.
		out, err := exec.Command("tcmd-handler", res.Method).Output()
		if err != nil {
			errorMarshalled, _ := json.Marshal(definitions.TcmdError{
				Error: err.Error(),
				Code:  definitions.HandleErrorCode,
			})
			c.Write(errorMarshalled)
			return
		}
		if len(out) <= 1 {
			// Make sure to at least return one byte.
			out = append(out, '\n')
		}

		fmt.Println("🙊 writing", res.Method, "to", addr)
		c.Write([]byte(out))
	}
	c.Close()
	fmt.Println("❌ closed", addr)
}
