package main

import "fmt"

// Maior = Definição de uma função que recebe dois inteiros e retorna um inteiro
// como possui a primeira letra maiúscula pode ser exportada para outros
// pacotes.
func Maior(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Essa função não tem parãmetros e nem retorno, apenas exibe uma string, além
// disso não pode ser exportada, apenas usada dentro desse pacote.
func exibeOlaMundo() {
	fmt.Println("Hello world!")
}

// Essa funcao possui retorno múltiplo. Quando houverem mais de 1 retorno é
// obrigatório o uso de parênteses, informando o tipo de cada retorno(podem ser
// diferentes)
func executaAteASomaSerMaiorQue100(n int) (int, int) {
	soma := 0
	iteracoes := 0
	for soma < 100 {
		soma += n
		iteracoes++
	}
	return soma, iteracoes
}

// Isso se chama retorno nomeado de valores, serve para tornar o que a função
// retorna mais visível. Nota-se que o return não possui valor, chama-se retorno
// limpo, esse só deve ser usado em funções pequenas(auxiliares por exemplo), pois
// pode acabar dificultando a legibilidade em funções grandes
func executaAteASomaSerMaiorQue100Ex2(n int) (soma, iteracoes int) {
	soma = 0
	iteracoes = 0
	for soma < 100 {
		soma += n
		iteracoes++
	}
	return
}

// Função anônima, também chamada de Closure. Pode ser criado uma cadeia de funções
// uma retornando outra.
func calculaSoma() func(int) int {
	soma := 0
	return func(x int) int {
		soma += x
		return soma
	}
}

func main() {

	// Função com um retorno
	fmt.Println(Maior(10, 7))

	// Função sem retorno
	exibeOlaMundo()

	// Retorna dois valores que são atribuidos as variáveis de forma posicional
	soma, iteracoes := executaAteASomaSerMaiorQue100(10)
	fmt.Printf("Soma: %d\nIteracoes: %d\n", soma, iteracoes)

	// Caso eu não queria usar um deles posso atribuir o resultado
	// ao símbolo _, dessa forma aquele valor será descartado
	_, iteracoes2 := executaAteASomaSerMaiorQue100Ex2(12)
	fmt.Printf("Iteracoes: %d\n", iteracoes2)

	// O retorno da função CalculaSoma é outra função, logo calculo é uma função
	// podemos chamar calculo passando o parâmetro solicitado na declaração da função
	// anonima dentro de calculaSoma
	calculo := calculaSoma()
	calculo(10)

}
