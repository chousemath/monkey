package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/chousemath/monkey/lexer"
	"github.com/chousemath/monkey/token"
)

const PROMPT = ">> "

// Start initiates the repl from main.go
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		// read from input source until a new line is encountered
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)

		// take the new line, create a new instance of a lexer, print all tokens
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
