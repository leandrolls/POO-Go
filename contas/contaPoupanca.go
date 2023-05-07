package contas

import "alura/go/poo/clientes"

type ContaPoupanca struct{
	Titular clientes.Titular
	NumeroAgencia int 
	NumeroConta int
	Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar{
		c.saldo -= valorDoSaque
		return "saque realizado"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito > 0 
	if podeDepositar{
		c.saldo += valorDoDeposito
		return "Deposito realizado"
	} else {
		return "Deposito nao aconteceu"
	}
}

func (c *ContaPoupanca) Obtersaldo() float64 {
	return c.saldo
}