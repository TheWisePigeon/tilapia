package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var accumulator int
	for {
		fmt.Print(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		for _, c := range input {
			switch c {
			case 's':
				accumulator = accumulator * accumulator
			case 'i':
				accumulator++
			case 'd':
				accumulator--
			case 'o':
				fmt.Println(accumulator)
			}
			if accumulator < 0 || accumulator == 256 {
				accumulator = 0
			}
		}
	}
}
