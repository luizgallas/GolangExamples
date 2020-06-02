package main

import (
	"fmt"
	"math"
	"reflect"
)

// Interfaces guardam definições de métodos que implementam essa interface.
// Diferente de Java ou outras linguagens, não há uma palavra reservada implements
// ou algo explicito que diga que aquele método implementa uma interface, apenas é
// preciso que seja declarada a definição de método definida na interface
type geometria interface {
	area() float64
	perimetro() float64
}

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}

// É um método que pertece a struct retangulo e implementa a interface geometria
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// É um método que pertece a struct retangulo e implementa a interface geometria
func (r retangulo) perimetro() float64 {
	return 2*r.largura + 2*r.altura
}

// É um método que pertece a struct circulo e implementa a interface geometria
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

// É um método que pertece a struct circulo e implementa a interface geometria
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Esse método recebe uma interface como parâmetro, ou seja, pode ser passado
// qualquer struct que implemente essa interface que irá funcionar, chamando suas
// funções relativas
func medidas(g geometria) {
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	r1 := retangulo{largura: 5.2, altura: 4.2}
	c1 := circulo{raio: 2.14}
	// Chamada de métodos que pertencem a struct é feito de forma diferente de
	// métodos tradicionais
	fmt.Println(r1.area())
	fmt.Println(r1.perimetro())

	// É possível usar a mesma declaração para circulo e retangulo implementando
	// soluções distintas para cada caso
	fmt.Println(c1.area())
	fmt.Println(c1.perimetro())

	// Chama a função medidas passando a struct retangulo, que implementa a interface
	// geometria
	medidas(r1)
	// Chama a função medidas passando a struct circulo, que implementa a interface
	// geometria
	medidas(c1)

	// Para cada um dos casso a interface irá verificar o tipo da struct e chamar seus
	// métodos relativos

}
