package products

import "fmt"

type Service interface{
	GetAll() ([]Product)
	GetById(id int) (Product,error)
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaPublicacion string) (Product,error)
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service)GetAll() []Product{
	return s.repository.GetAll()
}

func (s *service)GetById(id int) (Product,error){
	return s.repository.GetById(id)
}

func (s *service) Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error){
	if nombre==""{
		return Product{},fmt.Errorf("el campo nombre es requerido")
	}
	if color==""{
		return Product{},fmt.Errorf("el campo color es requerido")
	}
	if precio==0.0{
		return Product{},fmt.Errorf("el campo precio es requerido")
	}
	if stock==0{
		return Product{},fmt.Errorf("el campo stock es requerido")
	}
	if codigo==""{
		return Product{},fmt.Errorf("el campo codigo es requerido")
	}
	if fechaCreacion==""{
		return Product{},fmt.Errorf("el campo fechaPublicacion es requerido")
	}
	id := s.repository.LastID()
	id++
	return s.repository.Store(id,nombre,color,precio,stock,codigo,publicado,fechaCreacion)
}