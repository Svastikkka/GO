package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// taking input as line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()
	lowerString := strings.ToLower(inputString)
	formatString := strings.ReplaceAll(lowerString, " ", "")
	if strings.Contains(formatString, "a") && strings.HasPrefix(formatString, "i") && strings.HasSuffix(formatString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
