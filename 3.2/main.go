package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	divideNumbers(s)
}

func divideNumbers(s string) {
	s = strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), ",", ".")
	newString := strings.Split(s, ";")
	f1, err1 := strconv.ParseFloat(newString[0], 64)
	if err1 != nil {
		panic(err1)
	}
	f2, err2 := strconv.ParseFloat(newString[1], 64)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%.4f", float32(f1/f2))
}
