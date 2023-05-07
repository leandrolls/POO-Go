package contas

import "alura/go/poo/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar{
		c.saldo -= valorDoSaque
		return "saque realizado"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
	podeDepositar := valorDoDeposito > 0 
	if podeDepositar{
		c.saldo += valorDoDeposito
		return "Deposito realizado"
	} else {
		return "Deposito nao aconteceu"
	}
}


func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) string {
	podeTransferir := valorDaTransferencia < c.saldo
	if podeTransferir{
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return "Transferencia realizado"
	} else {
		return "Transferencia nao aconteceu"
	}
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}