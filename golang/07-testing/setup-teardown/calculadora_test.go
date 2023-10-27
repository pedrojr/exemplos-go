package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	fmt.Println("SetUp...")

	return func(t *testing.T) {
		fmt.Println("TearDown...")
	}
}

func TestSoma(t *testing.T) {
	tearDown := setupTestCase(t)
	defer tearDown(t)

	resultado := Soma(10, 2)
	assert.Equal(t, 12, resultado)
	fmt.Println("TestSoma")
}

func TestSubtracao(t *testing.T) {
	tearDown := setupTestCase(t)
	defer tearDown(t)

	resultado := Subtracao(10, 2)
	assert.Equal(t, 8, resultado)
	fmt.Println("TestSubtracao")
}
