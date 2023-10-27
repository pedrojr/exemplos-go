package calculadora

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := Soma(10, 2)
	assert.Equal(t, 12, resultado)
}

func TestSubtracao(t *testing.T) {
	resultado := Subtracao(10, 2)
	assert.Equal(t, 8, resultado)
}

func TestMultiplicacao(t *testing.T) {
	resultado := Multiplicacao(10, 2)
	assert.Equal(t, 20, resultado)
}

func TestDivisao(t *testing.T) {
	resultado := Divisao(10, 2)
	assert.Equal(t, 5, resultado)
}
