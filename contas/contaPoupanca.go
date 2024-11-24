package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	if valorDoSaque > c.saldo {
		return "Saldo insuficiente"
	}

	if valorDoSaque < 0 {
		return "Valor do saque menor que zero"
	}

	c.saldo -= valorDoSaque
	return "Saque realizado com sucesso"
}

func (c *ContaPoupanca) Depositar(valorDeDeposito float64) (string, float64) {
	if valorDeDeposito < 0 {
		return "Depósito menor que zero", c.saldo
	}

	c.saldo += valorDeDeposito
	return "Depósito realizado com sucesso", c.saldo
}