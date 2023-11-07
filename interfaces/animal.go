package interfaces

type Animal interface {
	respirar()
	comer()
	EsCarnivoro()
	EsComestible()
	tipo() string
	EstaVivo() bool
}
