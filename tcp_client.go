package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//Esto es para pedir info por teclado
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Dime host:port")
		return
		//Lo que hemos hecho arriba es que con ":" nos separa los datos
		//por teclado para que podamos pillarlos
	}
	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return

		}
	}
}

/*
   This file creates the main package, which declares the main() function. The function will use the imported packages to create a TCP client.
   The main() function gathers command line arguments in the arguments variable and makes sure that a value for host:port was sent.
   The CONNECT variable stores the value of arguments[1]to be used in the net.Dial() call.
   A call to net.Dial() begins the implementation of the TCP client and will connect you to the desired TCP server. The second parameter of net.Dial() has two parts; the first is the hostname or the IP address of the TCP server and the second is the port number the TCP server listens on.
   bufio.NewReader(os.Stdin) and ReadString() is used to read user input. Any user input is sent to the TCP server over the network using Fprintf().
   bufio reader and the bufio.NewReader(c).ReadString('\n') statement read the TCP server’s response. The error variable is ignored here for simplicity.
   The entire for loop that is used to read user input will only terminate when you send the STOP command to the TCP server.
*/
