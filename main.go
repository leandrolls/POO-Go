package main

import ("fmt"
		"alura/go/poo/contas"
)


func PagarBoleto(conta verificarConta, valorDoBoleto float64){
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface{
	Sacar(valor float64) string
}

func main() {


	contaDaSilvia := contas.ContaPoupanca{}
	contaDaSilvia.Depositar(200)

	fmt.Println(contaDaSilvia.Obtersaldo())
}