package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRepository struct{
	getByIdUsed bool
}

func (s *mockRepository) GetAll() ([]Product,error){
	publicado := true
	return []Product{{1,"lapicera","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},{2,"carpeta","rojo", 500.00, 90, "jklsfnvjfs", &publicado, "13/10/22"}},nil
}
func  (s *mockRepository) GetById(id int) (Product,error){
	s.getByIdUsed = true
	publicado := true
	return Product{1,"before update","azul", 100.00, 130, "jwnovinsjnfi", &publicado, "20/10/22"},nil
}
func (s *mockRepository) Store(id int,nombre, color string, precio float64, stock int, codigo string, publicado *bool, fechaCreacion string) (Product,error){
	return Product{},nil
}
func (s *mockRepository) LastID() (int,error){
	return 0,nil
}
func (s *mockRepository) ModifyAll(Product) (Product,error){
	return Product{},nil
}
func (s *mockRepository) Delete(id int) (Product,error){
	return Product{}, nil
}
func (s *mockRepository) ModifyValues(id int, nombre string, precio float64) (Product,error){
	product, err := s.GetById(1)
	if err != nil {
		return Product{}, err
	}
	product.Nombre = nombre
	return product, nil
}


func TestModifyValues(t *testing.T) {
	repository:= mockRepository{false}
	service := NewService(&repository)

	esperado := "after update"

	resultado, _ := service.ModifyValues(1,"after update", 300.00)

	assert.Equal(t,esperado,resultado.Nombre,"deben ser iguales")
	assert.True(t,repository.getByIdUsed,"debe ser utilizado")
}
