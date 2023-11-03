package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	paises["Chile"] = "Santiago"

	fmt.Println(paises)
	fmt.Println(paises["Chile"])
	fmt.Println(paises["-"])

	campeonato := map[string]int{
		"Barcelona":            34,
		"Universidad de Chile": 40,
		"Real Madrid":          38,
		"Manchester City":      39,
		"River Plate":          35,
	}
	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf(" el equipo %s tiene %d puntos \n", equipo, puntaje)

	}*/

	delete(campeonato, "River Plate")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Universidad de Chile"]
	fmt.Printf(" el puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)

}
