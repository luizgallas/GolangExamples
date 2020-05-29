// pacote principal de um programa em Go
package main

// importação consignada
import (
	"Golang_Example_Codes/examples"
	"fmt"
)

// função principal a ser chamada ao executar o programa
func main() {
	// Chama a função Hello do pacote examples
	examples.Hello()

	// Chama a função Print do pacote fmt dos pacotes padrões do Go
	fmt.Print(" World!\n")
}
