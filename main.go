package main

import (
	"flag"
	"fmt"
)

type Token string

const (
	I Token = "i"
	D Token = "d"
	S Token = "s"
	O Token = "o"
)

var accumulator = 0

func parse() {

}

func interpret() {

}

func main() {
	interactiveMode := flag.Bool("i", false, "Run the interpreter in interactive mode or not")
	flag.Parse()
	if *interactiveMode {
		fmt.Println("we doin it inter")
		return
	}
	fmt.Println("read from file")
}
