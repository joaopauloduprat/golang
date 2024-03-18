package main

import (
	"fmt"

	"github.com/joaopauloduprat/golang/testdocs/cachorro"
)

func main() {
	fmt.Println("Meu cachorro tem: ", cachorro.Idade(12), " anos")
}