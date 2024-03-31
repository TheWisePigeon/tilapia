package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Token string

const (
	I Token = "i"
	D Token = "d"
	S Token = "s"
	O Token = "o"
)

var accumulator int = 0
var tokens []Token

func tokenize(input string) {
	input = strings.TrimSuffix(input, "\n")
	for pos, r := range input {
		token := Token(r)
		switch token {
		case I, D, S, O:
			tokens = append(tokens, token)
		default:
			fmt.Printf("Illegal character: %s encountered in input at position %d\n", token, pos+1)
			tokens = []Token{}
		}
	}
}

func resetAccumulator() {
	if accumulator == -1 || accumulator == 256 {
		accumulator = 0
	}
}

func interpret() {
	for _, token := range tokens {
		switch token {
		case I:
			accumulator = accumulator + 1
			resetAccumulator()
		case D:
			accumulator = accumulator - 1
			resetAccumulator()
		case S:
			accumulator = accumulator * accumulator
			resetAccumulator()
		case O:
			fmt.Println(accumulator)
		default:
			panic(fmt.Sprintf("Invalid token %s", token))
		}
	}
	tokens = []Token{}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		tokenize(input)
		interpret()
	}
}
