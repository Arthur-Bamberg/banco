package main

import (
	"fmt"
	"banco/contas"
	"banco/clientes"
)

func PagarBoleto(conta VerificarConta, valorBoleto float64) string {
	return conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteArthur := clientes.Titular{Nome: "Arthur", CPF: "016.347.290.48", Profissao: "Desenvolvedor Backend"}
	contaDoArthur := contas.ContaCorrente{Titular: clienteArthur, NumeroAgencia: 589, NumeroConta: 123456}
	contaDoArthur.Depositar(300)

	clienteBruna := clientes.Titular{Nome: "Bruna", CPF: "016.347.290.48", Profissao: "Desenvolvedora Frontend"}
	contaDaBruna := contas.ContaCorrente{Titular: clienteBruna, NumeroAgencia: 589, NumeroConta: 123457}
	contaDaBruna.Depositar(100)

	fmt.Println(contaDoArthur.ObterSaldo())
	fmt.Println(contaDaBruna.ObterSaldo())
	fmt.Println(contaDoArthur.Transferir(200, &contaDaBruna))
	fmt.Println(contaDoArthur.ObterSaldo())
	fmt.Println(contaDaBruna.ObterSaldo())

	fmt.Println(PagarBoleto(&contaDoArthur, 50))

	fmt.Println(contaDoArthur.ObterSaldo())
}
