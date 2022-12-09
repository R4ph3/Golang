package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Por favor dime el numero del puerto")
		return
	}
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Saliendo del servidor TCP")
			return
		}

		fmt.Print("--> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))

	}

}

/*
   This file creates the main package, which declares the main() function. The function will use the imported packages to create a TCP server.
   The main() function gathers command line arguments in the arguments variable and includes error handling.
   The net.Listen() function makes the program a TCP server. This functions returns a Listener variable, which is a generic network listener for stream-oriented protocols.
   It is only after a successful call to Accept() that the TCP server can begin to interact with TCP clients.
   The current implementation of the TCP server can only serve the first TCP client that connects to it, because the Accept() call is outside of the for loop. In the Create a Concurrent TCP Server section of this guide, you will see a TCP server implementation that can serve multiple TCP clients using Goroutines.
   The TCP server uses regular File I/O functions to interact with TCP clients. This interaction takes place inside the for loop. Similarly to the TCP client, when the TCP server receives the STOP command from the TCP client, it will terminate.

*/
