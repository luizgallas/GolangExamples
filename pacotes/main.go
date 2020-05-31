package main

// Pacote "main" é o pacote principal de todo programa em Go, que gera um executável

// Importação consignada, onde os imports são envolvidos dentro dos parenteses.
// A quebra de linha nos pareteses é obrigatória
import (
	// As importações são relativas a pasta "src" presente nas pastas
	// definidas na GOPATH, por isso devemos incluir o nome do projeto
	// na importação
	"GolangExamples/pacotes/hello"
	"fmt"
	"os"

	// Renomeação de pacotes ao importar
	so "os"
)

// A função "main" do pacote "main" é o ponto de entrada dos programas em Go
func main() {
	// Erro de compilação, pois não existe uma função World neste pacote
	//World()

	// Chamando funções exportadas do pacote "hello"
	hello.World()
	hello.Go()

	// Chamando a função Println do pacote "fmt"
	fmt.Println("Fim!")

	// Erro de compilação porque não existe um pacote com nome "os" importado,
	// pois renomeamos para "so"
	os.Exit(1)

	// Chamada correta da função Exit presente no pacote "os" que foi renomeado para "so"
	so.Exit(0)
}
