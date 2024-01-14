package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/perpetua1g0d/ATT/pkg"
	"os"
)

func main() {
	arg := flag.String("tool", "", "tool choice to execute")
	flag.Parse()

	switch parsedFlag := *arg; parsedFlag {
	case "escape_quotes":
		var s string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			s = pkg.ReplaceQuotesWithEscape(scanner.Text())
			fmt.Println("\nformatted string:\n" + s)
			break
		}
	default:
		return
	}
}
