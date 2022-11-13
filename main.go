package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nPlease enter a valid ISO country code...")
	scanner.Scan()
	countryCode := scanner.Text()

	fmt.Printf("You have chosen %s", countryCode)
}
