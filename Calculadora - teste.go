package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestMultiplica2(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 250

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestSubtrai(t *testing.T) {
	teste := subtrai(2, 10)
	resultado := 8

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestSubtrai2(t *testing.T) {
	teste := subtrai(2, 10)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestDivide(t *testing.T) {
	teste := divide(10)
	resultado := 10

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}

func TestDivide2(t *testing.T) {
	teste := divide(10)
	resultado := 8

	if teste != resultado {
		t.Error("Valor esperado: ", resultado, "Valor retornado: ", teste)
	}
}
