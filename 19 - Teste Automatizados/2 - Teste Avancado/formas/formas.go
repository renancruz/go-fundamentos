package formas

import (
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	Alutra  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Alutra * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}
