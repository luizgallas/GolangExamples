package main

import "fmt"

func main() {

	a, b := 10, 20
	sub := a - b

	// No if os parênteses são opcionais, porém a boa prática é não utilizar
	if sub < 0 {
		fmt.Println("Negativo")
	} else {
		fmt.Println("Positivo")
	}

	sexo := "M"
	// Expressão que queremos verificar
	switch sexo {
	case "M": // Verifica se sexo é igual a M
		fmt.Println("Masculino")
	case "F": // Verifica se sexo é igual a F
		fmt.Println("Feminino")
	default: // Caso não caia em nenhum dos anteriores vem para o default
		fmt.Println("Outro")
	}

}
