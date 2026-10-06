package main

import (
	"fmt"
	"reflect"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaBruno := ContaCorrente{
		titular:       "Bruno",
		numeroAgencia: 5542,
		numeroConta:   4521,
		saldo:         125.50,
	}

	contaBruna := ContaCorrente{
		"Bruna",
		222,
		11122,
		200.00,
	}

	fmt.Println(reflect.TypeOf(contaBruno))
	fmt.Println(contaBruna)
}
