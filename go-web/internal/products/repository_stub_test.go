package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubRepository struct{}

func (s stubRepository) GetAll() ([]Product,error){
	publicado := true
	return []Product{{1,"lapicera","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"}},nil
}
func  (s stubRepository) GetById(id int) (Product,error){
	return Product{},nil
}
func (s stubRepository) Store(id int,nombre, color string, precio float64, stock int, codigo string, publicado *bool, fechaCreacion string) (Product,error){
	return Product{},nil
}
func (s stubRepository) LastID() (int,error){
	return 0,nil
}
func (s stubRepository) ModifyAll(Product) (Product,error){
	return Product{},nil
}
func (s stubRepository) Delete(id int) (Product,error){
	return Product{}, nil
}
func (s stubRepository) ModifyValues(id int, nombre string, precio float64) (Product,error){
	return Product{}, nil
}

func TestGetAll(t *testing.T){
	repository := stubRepository{}
	service := NewService(repository)
	
	publicado := true
	esperado:=[]Product{{1,"lapicera","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"}}
	resultado,_ :=service.GetAll()

	assert.Equal(t,esperado,resultado , "deben ser iguales")
}