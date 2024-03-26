package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		text   string
		result []string
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = strings.ToLower(scanner.Text())
	charText := strings.Split(text, "")

	for i, char := range charText {
		var item string
		charText[i] = strings.ToUpper(char)

		item = strings.Join(charText, "")
		result = append(result, item)
		charText[i] = strings.ToLower(char)
	}

	fmt.Println(result)
}
