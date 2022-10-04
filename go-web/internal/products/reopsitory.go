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
	products := r.GetAll()
	var product Product
	for _,p:=range products{
		if p.ID == id{
			product = p
			break
		}
	}
	if product == (Product{}) {
		return Product{},fmt.Errorf("no existe el producto con id %d", id)
	}
	return product,nil
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