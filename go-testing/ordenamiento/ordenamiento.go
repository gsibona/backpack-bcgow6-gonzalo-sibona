package ordenamiento

import (
	"sort"
)

func ordenar(lista []int) []int{
	sort.Slice(lista, func(i, j int) bool {
		return lista[i] < lista[j]
	})
	return lista
}