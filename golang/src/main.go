package main

import "data_structures_and_algorithms/pkg/ds"

func main() {
	list := ds.BuildList([]int{1, 2, 3, 4, 5})
	list.Traverse()
}
