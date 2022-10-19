package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	lista := []int{5,2,7,4,11,1}
	esperado := []int{1,2,4,5,7,11}

	resultado := ordenar(lista)

	assert.Equal(t,esperado,resultado,"deben ser iguales")
}