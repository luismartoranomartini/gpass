package main

import (
	"crypto/rand"
	"math/big"
)

const (
	letras   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitos  = "0123456789"
	simbolos = "!@#$%^&*()-_=+[]{}<>?"
)

func alfabeto(comSimbolos bool) string {
	a := letras + digitos
	if comSimbolos {
		a += simbolos
	}
	return a
}

func indice(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, nil
	}
	return int(n.Int64()), nil
}

func gerar(n int, alfa string) (string, error) {
	saida := make([]byte, n)

	for i := range saida {
		j, err := indice(len(alfa))
		if err != nil {
			return "", err
		}
		saida[i] = alfa[j]
	}
	return string(saida), nil
}
