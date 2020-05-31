package main

import "fmt"

// Estruturas encapsulam atributos de qualquer tipo e podem ser exportadas para
// outros pacotes caso sua primeira letra seja maiuscula
type pessoa struct {
	nome  string
	idade int
	pai   *pessoa // É um ponteiro para outra pessoa, ou seja, guarda a refêrencia
} // para o endereço de memória onde está alocado a struct e seu valor

func main() {
	//Definição e atribuição de struct de forma nominada
	jose := pessoa{nome: "Jose", idade: 48, pai: nil}
	//Definição e atribuição de struct de forma posicional
	lucas := pessoa{"Lucas", 18, &jose}

	// Podemos pegar as informações de José através de Lucas visto que existe uma
	// referência à struct jose guardada no atributo pai
	fmt.Printf("%s tem %d anos e é filho de %s que tem %d anos\n", lucas.nome,
		lucas.idade, lucas.pai.nome, lucas.pai.idade)
}
