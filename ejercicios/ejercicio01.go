package ejercicios

import (
	"strconv"
)

func Ejercicio1(num string) (string, int) {

	numeroConvertido, err := strconv.Atoi(num)
	if err != nil {
		return "Hubo un error" + err.Error(), 0
	}

	if numeroConvertido > 100 {
		return "El numero es mayor  a 100", numeroConvertido
	} else if numeroConvertido == 100 {
		return "El numero es igual  a 100", numeroConvertido

	} else {
		return "El numero es menor  a 100", numeroConvertido

	}
}
