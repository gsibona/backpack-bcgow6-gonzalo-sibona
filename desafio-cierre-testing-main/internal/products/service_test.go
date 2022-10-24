package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubRepository struct{}

func (r stubRepository) GetAllBySeller(sellerID string) ([]Product, error){
	if sellerID == "1"{
		return []Product{{"1","1","descripcion",10.00},{"2","1","descripcion",15.00}},nil
	}
	return []Product{},errors.New("this is a error")
}

func TestGetAllBySeller(t *testing.T) {
	repository := stubRepository{}
	service := NewService(repository);
	
	esperado := []Product{{"1","1","descripcion",10.00},{"2","1","descripcion",15.00}}

	res,err:= service.GetAllBySeller("1")

	assert.Nil(t,err)
	assert.Equal(t,esperado,res)

	esperadoErr := "this is a error"

	res,err = service.GetAllBySeller("2")

	assert.Empty(t,res)
	assert.EqualError(t,err,esperadoErr)
}