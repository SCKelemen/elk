package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SCKelemen/elk/scanner"
	"github.com/SCKelemen/elk/token"
)

const PROMPT = "🦌> "

func Start(in io.Reader, out io.Writer) {
	reader := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		read := reader.Scan()
		if !read {
			return
		}

		line := reader.Text()
		s := scanner.New(line)

		for tok := s.NextToken(); tok.Kind != token.EOF; tok = s.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
