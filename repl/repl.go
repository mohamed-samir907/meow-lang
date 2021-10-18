package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mohamed-samir907/meow/lexer"
	"github.com/mohamed-samir907/meow/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		fmt.Printf(program.String() + "\n")

		/* for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		} */
	}
}
