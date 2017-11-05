package main

type Matrix struct {
	dimensionx int
	dimensiony int
	fields     [][]int
}

func NewMatrix(dimensionx int, dimensiony int) *Matrix {
	var m Matrix
	m.dimensionx = dimensionx
	m.dimensiony = dimensiony

	//init fields
	m.fields = make([][]int, dimensionx)
	for i := 0; i < dimensionx; i = i + 1 {
		m.fields[i] = make([]int, dimensiony)
	}

	return &m
}
