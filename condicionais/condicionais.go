package main

import "fmt"

func main() {

	// No if os parênteses são opcionais, porém a boa prática é não utilizar
	// A chave de abertura precisa vir ao lado da expressão
	// palavra reservada else ou else if precisa vir ao lado a chave de fechamento
	idade := 58
	if idade >= 60 {
		fmt.Println("Está no grupo de risco")
	} else {
		fmt.Println("Não está no grupo de risco")
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
