package main

// Pacote "main" é o pacote principal de todo programa em Go

// Importação consignada, onde os imports são envolvidos dentro dos parenteses.
// A quebra de linha nos pareteses é obrigatória
import (
	// As importações são relativas a pasta "src" presente nas pastas
	// definidas na GOPATH, por isso devemos incluir o nome do projeto
	// na importação
	"Golang_Example_Codes/pacotes/hello"
	"fmt"
)

// A função "main" do pacote "main" é o ponto de entrada dos programas em Go
func main() {

	// Chamando funções exportadas do pacote "hello"
	hello.World()
	hello.Go()

	// Chamando a função Println do pacote "fmt"
	fmt.Println("Fim!")
}