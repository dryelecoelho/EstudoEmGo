package main

import (
	"fmt"
)

//atribuicao dos valores as variáveis

var x int = 42
var y string = "james bond"
var z bool = true

func main() {

//atribuicao a uma única variável

	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)
}
