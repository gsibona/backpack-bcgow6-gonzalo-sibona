package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubDB struct{}

func (db stubDB) Read(data interface{}) error{
	publicado := true
	products := &[]Product{
		{1,"lapicera","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},
		{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"},
	}
	t,_ := data.(*[]Product)
	*t = *products
	return nil
}
func (db stubDB) Write(data interface{}) error{
	return nil
}


func TestGetAll(t *testing.T){
	db := stubDB{}
	repository := NewRepository(db)

	publicado:=true
	esperado:= []Product{
		{1,"lapicera","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},
		{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"},
	}

	respuesta,err := repository.GetAll()

	assert.Nil(t,err)
	assert.Equal(t,esperado,respuesta)
}