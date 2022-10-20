package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct{
	products []Product 
	readUsed bool
}

func (db *mockDB) Read(data interface{}) error{
	if db.products == nil{
		publicado := true
		db.products = []Product{
			{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},
			{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"},
		}
	}
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

	esperado := "after update"
	
	_, err:= repository.ModifyValues(1,"after update", 300.00)
	assert.Nil(t,err)
	
	resultado,err:= repository.GetById(1)
	assert.Nil(t,err)

	assert.Equal(t,esperado,resultado.Nombre,"deben ser iguales")
	assert.True(t,db.readUsed,"debe ser utilizado")
}
