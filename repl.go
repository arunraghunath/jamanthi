package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/user"

	"github.com/arunraghunath/jamanthi/token"

	"github.com/arunraghunath/jamanthi/lexer"
)

const PROMPT = ">>"

func StartInteractive(r io.Reader, w io.Writer) {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Hi %s!!, This is the interactive terminal for the jamanthi programming language.\n",
		user.Username)
	fmt.Printf(PROMPT)

	for {
		scanner := bufio.NewScanner(r)
		if scanned := scanner.Scan(); !scanned {
			return
		} else {
			l := lexer.New(scanner.Text())
			for tok := l.NextToken(); tok.Typ != token.EOF; tok = l.NextToken() {
				fmt.Printf("Token details are Type ==> %d, Content ==> %s\n",
					tok.Typ, tok.Val)
			}
			fmt.Printf(PROMPT)
		}
	}

}
