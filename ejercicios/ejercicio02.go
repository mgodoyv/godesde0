package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numeroMult int
var err error

func TablaMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero para presentar su tabla de multiplicar ")
		if scanner.Scan() {
			numeroMult, err = strconv.Atoi(scanner.Text())
		}
		if err != nil {
			continue
		} else {
			break
		}

	}

	for i := 1; i <= 10; i++ {
		/*fmt.Print("El resultado de la multiplicacion de " + strconv.Itoa(numeroMult) + "*" + strconv.Itoa(i) + " es ")
		fmt.Println(numeroMult * i)
		*/
		fmt.Printf("el resultado de %d x %d = %d \n", numeroMult, i, numeroMult*i)

	}

}
