package products

import (
	"fmt"
)

type Product struct{
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

type Repository interface{
	GetAll() ([]Product)
	GetById(id int) (Product,error)
	Store(id int,nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error)
	LastID() (int)
	ModifyAll(Product) (Product,error)
	Delete(id int) (Product,error)
	ModifyValues(id int, nombre string, precio float64) (Product,error)
}

type repository struct{}

var ps []Product

var lastId int

func NewRepository () Repository{
	return &repository{}
}

func (r *repository) GetAll() [] Product{
	return ps
}

func (r *repository) GetById(id int) (Product,error){
	for _,p:=range ps{
		if p.ID == id{
			return p,nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", id)	
}

func (r *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error){
	p := Product{id,nombre,color,precio,stock,codigo,publicado,fechaCreacion}
	ps = append(ps,p)
	lastId = p.ID
	return p,nil
}

func (r *repository) LastID() int{
	return lastId
}

func (r *repository) ModifyAll(product Product) (Product,error){
	for i,p:=range ps{
		if p.ID == product.ID{
			ps[i] = product
			return ps[i],nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", product.ID)
}

func (r *repository) Delete(id int) (Product,error){
	for i,p:=range ps{
		if p.ID == id{
			ps = append(ps[:i], ps[i+1:]...)
			return p,nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", id)
}

func(r *repository) ModifyValues(id int, nombre string, precio float64) (Product,error){
	for i,p:=range ps{
		if p.ID == id{
			ps[i].Nombre = nombre
			ps[i].Precio = precio
			return ps[i],nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", id)
}