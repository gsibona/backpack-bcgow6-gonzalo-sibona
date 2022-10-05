package products

import "fmt"

type Service interface{
	GetAll() ([]Product)
	GetById(id int) (Product,error)
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error)
	ModifyAll(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error)
	Delete(id int) (Product, error)
	ModifyValues(id int, nombre string, precio float64) (Product,error)
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
	err := notNull(nombre,color,precio,stock,codigo,fechaCreacion)
	if err!=nil{
		return Product{},err
	}
	id := s.repository.LastID()
	id++
	return s.repository.Store(id,nombre,color,precio,stock,codigo,publicado,fechaCreacion)
}

func (s *service) ModifyAll(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error){
	err := notNull(nombre,color,precio,stock,codigo,fechaCreacion)
	if err!=nil{
		return Product{},err
	}
	return s.repository.ModifyAll(Product{id,nombre,color,precio,stock,codigo,publicado,fechaCreacion})
}



func notNull(nombre, color string, precio float64, stock int, codigo string, fechaCreacion string) (err error){
	if nombre==""{
		err = fmt.Errorf("el campo nombre es requerido")
	}
	if color==""{
		err = fmt.Errorf("el campo color es requerido")
	}
	if precio==0.0{
		err = fmt.Errorf("el campo precio es requerido")
	}
	if stock==0{
		err = fmt.Errorf("el campo stock es requerido")
	}
	if codigo==""{
		err = fmt.Errorf("el campo codigo es requerido")
	}
	if fechaCreacion==""{
		err = fmt.Errorf("el campo fechaPublicacion es requerido")
	}
	return
}

func (s *service) Delete(id int) (Product,error){
	return s.repository.Delete(id)
}

func (s *service) ModifyValues(id int, nombre string, precio float64) (Product,error){
	err := notNull(nombre,"nil",precio,1,"nil","nil")
	if err!=nil{
		return Product{},err
	}
	return s.repository.ModifyValues(id,nombre,precio)
}