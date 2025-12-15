package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) > 0 {
			first := strings.ToLower(words[0])
			fmt.Printf("Your command was: %s\n", first)
		}
	}
}
