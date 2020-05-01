package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Birl")
	esperado := "Olá, Birl"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
