package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)

	if a*b%2 == 0 {
		fmt.Printf("%s", "Even")
	} else {
		fmt.Printf("%s", "Odd")
	}
}
