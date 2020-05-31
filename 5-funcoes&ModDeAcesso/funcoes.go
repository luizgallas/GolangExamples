package funcoes_e_mod_de_acesso

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
