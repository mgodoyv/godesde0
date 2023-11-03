package arreglosslices

import "fmt"

var tabla [10]int = [10]int{190, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)

	tabla2 := [10]string{"Mauro", "Mati"}
	fmt.Println(tabla2)

	for i := 1; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
