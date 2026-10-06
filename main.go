package main

import (
	"fmt"
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

	fmt.Println(contaBruno)
	fmt.Println(contaBruna)

	var contaCris *ContaCorrente
	contaCris = new(ContaCorrente)
	contaCris.titular = "Cris"

	fmt.Println(*contaCris)
}
