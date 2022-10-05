package products

import (
	"fmt"

	"github.com/gsibona/backpack-bcgow6-gonzalo-sibona/go-web/pkg/store"
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
	GetAll() ([]Product,error)
	GetById(id int) (Product,error)
	Store(id int,nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error)
	LastID() (int,error)
	ModifyAll(Product) (Product,error)
	Delete(id int) (Product,error)
	ModifyValues(id int, nombre string, precio float64) (Product,error)
}

type repository struct{
	db store.Store
}


func NewRepository(db store.Store) Repository{
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product,error){
	var ps []Product
	if err:=r.db.Read(&ps); err!=nil{
		return []Product{},err
	}
	return ps,nil
}

func (r *repository) GetById(id int) (Product,error){
	var ps []Product
	r.db.Read(&ps)
	for _,p:=range ps{
		if p.ID == id{
			return p,nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", id)	
}

func (r *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Product,error){
	p := Product{id,nombre,color,precio,stock,codigo,publicado,fechaCreacion}
	var ps []Product
	r.db.Read(&ps)
	ps = append(ps,p)
	if err:=r.db.Write(ps); err!=nil {
		return Product{},err
	}
	return p,nil
}

func (r *repository) LastID() (int,error){
	var ps []Product
	if err:=r.db.Read(&ps); err!=nil{
		return 0,err
	}
	if len(ps)==0{
		return 0,nil
	}
	return ps[len(ps)-1].ID,nil
}

func (r *repository) ModifyAll(product Product) (Product,error){
	var ps []Product
	r.db.Read(&ps)
	for i,p:=range ps{
		if p.ID == product.ID{
			ps[i] = product
			if err:=r.db.Write(ps); err!=nil {
				return Product{},err
			}
			return ps[i],nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", product.ID)
}

func (r *repository) Delete(id int) (Product,error){
	var ps []Product
	r.db.Read(&ps)
	for i,p:=range ps{
		if p.ID == id{
			ps = append(ps[:i], ps[i+1:]...)
			if err:=r.db.Write(ps); err!=nil {
				return Product{},err
			}
			return p,nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", id)
}

func(r *repository) ModifyValues(id int, nombre string, precio float64) (Product,error){
	var ps []Product
	r.db.Read(&ps)
	for i,p:=range ps{
		if p.ID == id{
			ps[i].Nombre = nombre
			ps[i].Precio = precio
			if err:=r.db.Write(ps); err!=nil {
				return Product{},err
			}
			return ps[i],nil
		}
	}
	return Product{},fmt.Errorf("no existe el producto con id %d", id)
}