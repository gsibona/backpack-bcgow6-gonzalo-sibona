package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct{
	products []Product 
	readUsed bool
}

func (db *mockDB) Read(data interface{}) error{
	t,_ := data.(*[]Product)
	*t = db.products
	db.readUsed = true
	return nil
}
func (db *mockDB) Write(data interface{}) error{
	db.products = data.([]Product)
	return nil
}

func TestModifyValues(t *testing.T) {
	db := mockDB{}
	repository:= NewRepository(&db)
	publicado := true
	db.products = []Product{
		{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},
		{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"},
	}

	esperado := "after update"
	
	_, err:= repository.ModifyValues(1,"after update", 300.00)
	assert.Nil(t,err)
	
	resultado,err:= repository.GetById(1)
	assert.Nil(t,err)

	assert.Equal(t,esperado,resultado.Nombre,"deben ser iguales")
	assert.True(t,db.readUsed,"debe ser utilizado")

	esperadoError := fmt.Errorf("no existe el producto con id %d", 3)
	
	res2,err:= repository.ModifyValues(3,"after update", 300.00)

	assert.Empty(t,res2)
	assert.EqualError(t,esperadoError,err.Error())
}

func TestGetById(t *testing.T) {
	db := mockDB{}
	repository:= NewRepository(&db)
	publicado := true
	db.products = []Product{
		{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},
		{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"},
	}

	esperado := fmt.Errorf("no existe el producto con id %d", 3)

	respuesta,err := repository.GetById(3)
	
	assert.Empty(t,respuesta)
	assert.EqualError(t,esperado ,err.Error())
}


func TestStore(t *testing.T) {
	db := mockDB{}
	repository:= NewRepository(&db)
	publicado := true
	db.products = []Product{
		{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},
		{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"},
	}

	esperado := Product{3,"libro","naranja", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"}

	resultado1,err := repository.Store(3,"libro","naranja", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22")
	assert.Nil(t,err)

	resultado2,err:= repository.GetById(3)
	assert.Nil(t,err)

	assert.Equal(t,esperado,resultado1)
	assert.Equal(t,esperado,resultado2)
}