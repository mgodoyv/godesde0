package funciones

import (
	"fmt"
	"strconv"
)

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClousure() {
	tabladel := 2
	miTabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(miTabla())
	}
}

func TablasHastaX(x int) {
	for j := 1; j <= x; j++ {
		fmt.Println("Tabla del " + strconv.Itoa(j))

		for i := 1; i <= 10; i++ {
			tablaX := tabla(i)
			res := tablaX() * j
			fmt.Println(res)
		}
	}
}
