package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stdin = bufio.NewScanner(os.Stdin)

func scanner() (s string) {
	if stdin.Scan() {
		s = strings.TrimSpace(stdin.Text())
	} else {
		log.Fatalln(stdin.Err())
	}

	return
}

func convToInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

func main() {
	i1 := scanner()
	i23 := strings.Split(scanner(), " ")
	s := scanner()

	fmt.Printf("%d %s", convToInt(i1)+convToInt(i23[0])+convToInt(i23[1]), s)
}
