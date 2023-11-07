package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
	padre      string
}

func (h *Hombre) Respirar()          { h.respirando = true }
func (h *Hombre) Comer()             { h.comiendo = true }
func (h *Hombre) Pensar()            { h.pensando = true }
func (h *Hombre) Sexo() string       { return "Hombre" }
func (h *Hombre) Paternidad() string { return "No" }
func (h *Hombre) EstaVivo() bool     { return true }

func (men *Hombre) AddHombre(edad int, altura float32, peso float32) {
	men.edad = edad
	men.altura = altura
	men.peso = peso

}
