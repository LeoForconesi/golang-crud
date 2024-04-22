package services

type coso interface {
	// Quicksort ordena un arreglo de menor a mayor
	QuickSort(array []int) []int

	// binary search es un algoritmo de busqueda
	BinarySearch(element int) bool
}

type Algorithms struct {
	AElement int
}

func (ae Algorithms) BinarySearch()
