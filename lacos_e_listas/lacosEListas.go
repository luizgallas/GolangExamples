package main

import "fmt"

func main() {

	// Laço for tradicional
	// Não pode ser usado parênteses, chave de abertura vai do lado da expressão
	for i := 0; i < 10; i++ {
		fmt.Println(i * 10)
	}

	//Matrizes = listas com tamanho fixo
	// Sintaxe [tamanho]tipo{conteudo}
	// Não suportam tipos distintos
	matriz := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Slices = listas com tamanho dinâmico
	// Sintaxe []tipo{conteudo}
	// Não suportam tipos distintos
	slice := []string{"Olá", "Mundo", "!"}

	// Laço for range = laço forEach
	// 1º Indice, 2º valor atual, 3º lista a ser iterada
	// Para cada inteiro dentro da lista matriz soma e exibe
	// Operador _ significa que queremos ignorar algum dos retornos de uma função que contém retorno multíplo
	var soma int
	for _, inteiro := range matriz {
		soma += inteiro
		fmt.Println(soma)
	}

	for indice, palavra := range slice {
		fmt.Println(palavra, indice)
	}

	// Ao adicionar novos valores em uma slice(ou matriz) é gerada uma nova lista contendo os valores
	// da lista original e os novos valores
	// Sintaxe: primeiro parâmetro é uma slice do tipo T, todos os parâmetros seguintes são valores
	// a serem adicionados na lista
	// Caso a lista seja uma matriz: Se a matriz anterior de s for muito pequena para caber todos os
	// valores uma matriz gigante será alocada. A slice retornada apontará para a nova matriz alocada.
	novaSlice := append(slice, "Adicionando", "novos", "valores", "na", "slice")

	for _, palavra := range novaSlice {
		fmt.Println(palavra)
	}
}
