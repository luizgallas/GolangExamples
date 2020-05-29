package main

import (
	"fmt"
)

// Constantes
const pi = 3.14

// podem ser usadas com consignado para representar enumerações
const (
	segunda = "Segunda-feira"
	terca   = "Terã-feira"
	quarta  = "Quarta-feira"
	quinta  = "Quinta-feira"
	sexta   = "Sexta-feira"
	sabado  = "Sábado"
	domingo = "Domingo"
)

func main() {

	// Declaração de variável com atribuição de valor
	var nome string = "Luiz"

	// Declaração de variável sem atribuição de valor, recebe valor default 0
	var idade int

	// Declaração de variável sem atribuição de valor, recebe valor default false
	var maiorDeIdade bool

	// Declaração de variável sem atribuição de valor, recebe valor default 0
	var altura float32

	// Declaração de variável sem atribuição de valor, recebe valor default 0
	var inteiroPositivo uint

	// Declaração de variável sem atribuição de valor, recebe valor default 0
	var complexos complex128 // Representa numeros complexos matemáticos

	// Declaração de variável sem atribuição de valor, recebe valor default 0
	var unicode rune // pseudônimo para int32 que representa um ponto de código Unicode

	// Declaração de variável sem atribuição de valor, recebe valor default 0
	var short byte // pseudônimo para int8

	fmt.Println(nome, idade, maiorDeIdade, altura, inteiroPositivo, complexos, unicode, short)

	// Declara e atribui um valor de forma curta
	sobrenome := "Gallas"

	// Declaração e atribuição omitindo o tipo
	var omitindoTipo = 10

	fmt.Println(sobrenome, omitindoTipo, pi)

	// É possível declarar variáveis em conjunto
	a, b := 1, 2

	// Declaração e atribuição consignada
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

}
