package dividir

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num,den:= 17,0
	_,err:= Dividir(num,den)
	assert.Errorf(t,errors.New("El denominador no puede ser 0"),err.Error())
}