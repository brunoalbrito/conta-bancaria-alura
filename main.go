package main

import "fmt"

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

	fmt.Println(contaBruno)
}
