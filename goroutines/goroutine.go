package main

import (
	"fmt"
	"time"
)

// CursorType uma simples struct para armazenar o ponteiro de texto
// a ser exibido, e a posição atual deste ponteiro
type CursorType struct {
	Value    string
	Position int
	Writing bool
}

// Algumas constantes para definir tempos de animações como
// o tempo que o cursor leva para piscar, o tempo para simular
// a escrita, etc...
const (
	blinkRate = 400  // animação do cursor
	writeTime = 56   // animação de escrita
	eraseTime = 16   // animação apagar
	waitTime  = 3000 // espera após exibir todo o texto
)

// text é o texto a ser exibido na animação
const text = `
package main

import "fmt"

func main() {
    fmt.Println("Olá! Obrigado por participar deste Minicurso ;)")
}
`

func main() {
	// Converte o texto simples para um slice de rune (Unicode)
	runeText := []rune(text)

	// cria uma instancia do CursorType para armazenar informações do cursor
	cursor := &CursorType{"|", 0, false}

	// cria um canal para atualizar o terminal sempre que necessário
	update := make(chan bool)

	// cria uma goroutine de uma função anônima, com um laço infinito,
	// que será executada infinitamente (até o término do programa),
	// atualizando o valor do cursor de acordo com o tempo especificado
	// na constante blinkRate, através do time.Sleep
	go func(c *CursorType) {
		// cria um laço infinito
		for {
			if c.Value == "|" {
				if !c.Writing {
				    c.Value = " "
				}
			} else {
				c.Value = "|"
			}
			
			if !c.Writing {
			    // direciona o valor true para o canal update apenas para atualizar o terminal
			    update <- true
			}

			// pausa a execução desta goroutine para que o cursor seja atualizado novamente
			// apenas no intervalo de tempo desejado
			time.Sleep(blinkRate * time.Millisecond)
		}
	}(cursor)

	// cria uma goroutine de uma função anônima, com um laço infinito,
	// que será executada infinitamente (até o término do programa),
	// aguardando pelo canal update, para atualizar o texto do terminal
	go func(c *CursorType) {
		// cria um laço infinito
		for {
			// aguarda o canal update
			<-update

			// limpa o console
			fmt.Print("\x0c")

			// exibe o texto atualizado, convertendo o slice de rune (Unicode) para string
			fmt.Printf("%s%s", string(runeText[:cursor.Position]), c.Value)
		}
	}(cursor)

	// cria um loop infinito para alterar o texto a ser exibido
	for {
		// Se a posição do cursor chegou ao final do texto, define a flag writing para false
		// e chama o time.Sleep para aguardar o tempo de espera definido na constante waitTime.
		// Observe que estamos na goroutine principal
		if cursor.Position == len(runeText) {
			cursor.Writing = false
			time.Sleep(waitTime * time.Millisecond)
		} else if cursor.Position == 0 {
			// Se o cursor chegou ao início, define a flag writing como true
			cursor.Writing = true
		}

		// Se estiver escrevendo, incrementa a posição do cursor e aguarda o tempo de escrita
		if cursor.Writing {
			cursor.Position++
			time.Sleep(writeTime * time.Millisecond)
		} else if cursor.Position > 0 {
			// Se estiver apagando, decrementa o cursor e aguarda o tempo de apagar
			cursor.Position--
			time.Sleep(eraseTime * time.Millisecond)
		}

		// Envia um valor ao canal update para atualizar o texto do terminal
		update <- true
	}
}
