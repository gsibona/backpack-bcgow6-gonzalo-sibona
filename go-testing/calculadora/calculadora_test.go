package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta (t *testing.T){
	num1, num2 :=5,3
	esperado := 2

	resultado := resta(num1,num2)

	assert.Equal(t, resultado,esperado,"deben ser iguales")

}