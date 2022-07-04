package formas

import (
	"math"
	"testing"
)

type cenarioRetangulo struct {
	Retangulo
	areaEsperada float64
}

type cenarioCirculo struct {
	Circulo
	areaEsperada float64
}

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

		// Exemplo utilizando cenario
		ret2 := cenarioRetangulo{Retangulo{10, 12}, 120}
		areaRecebida2 := ret2.Area()

		if ret2.areaEsperada != areaRecebida2 {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida2, ret2.areaEsperada)
		}

	})

	t.Run("Ciruclo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

}
