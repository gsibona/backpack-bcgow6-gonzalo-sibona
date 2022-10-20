package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifyAll(t *testing.T) {
	publicado:=true
	db := mockDB{
		products: []Product{{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"}},
	}
	repository := NewRepository(&db)
	service := NewService(repository)

	esperado := Product{1,"after update","verde", 120.00, 120, "jwnovinsjnfi", &publicado, "20/10/22"}

	resultado, err:= service.ModifyAll(1,"after update", "verde", 120.00,120,"jwnovinsjnfi",&publicado,"20/10/22")
	assert.Nil(t,err)

	assert.Equal(t,esperado,resultado)
	assert.True(t,db.readUsed)
}

func TestDelete(t *testing.T) {
	publicado:=true
	db := mockDB{
		products: []Product{{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"}},
	}
	repository := NewRepository(&db)
	service := NewService(repository)

	esperado1 := Product{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"}
	esperado2 := "no existe el producto con id 2"

	res1,err := service.Delete(1)
	assert.Nil(t,err)

	_,err = service.Delete(2)

	assert.EqualError(t,err,esperado2)
	assert.Equal(t,esperado1,res1)
}