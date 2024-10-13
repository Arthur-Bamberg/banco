package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	if valorDoSaque > c.saldo {
		return "Saldo insuficiente", c.saldo
	}

	if valorDoSaque < 0 {
		return "Valor do saque menor que zero", c.saldo
	}

	c.saldo -= valorDoSaque
	return "Saque realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Depositar(valorDeDeposito float64) (string, float64) {
	if valorDeDeposito < 0 {
		return "Depósito menor que zero", c.saldo
	}

	c.saldo += valorDeDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (string, float64) {
	if valorDaTransferencia > c.saldo {
		return "Saldo insuficiente", c.saldo
	}

	if valorDaTransferencia < 0 {
		return "Valor da transferência menor que zero", c.saldo
	}

	c.saldo -= valorDaTransferencia
	contaDestino.saldo += valorDaTransferencia
	return "Transferência realizada com sucesso", c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
