package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func shell() {
	//Esto es para leer la entrada por teclado
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ") //Para poner una flechita donde irá el comando
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		//Esto lleva la ejecución del input en caso de que sea erroneo
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {

	//Eliminamos el \n de la terminal
	input = strings.TrimSuffix(input, "\n")
	//Preparamos el comando a ejecutar
	cmd := exec.Command(input)

	//Le decimos que pase por el sistema el output de lo que escribamos
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//Ejecuta el comando
	return cmd.Run()

}
