package main

import (
	"bufio"
	"os"
)

func Scan(input *string, scanner *bufio.Scanner) {
	scanner.Scan()
	*input = scanner.Text()
}

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	Scan(&input, scanner)
}
