package main

import (
	"fmt"
	"sort"
)




func fct(line []float64) {
	for _,v:= range line {
		fmt.Printf("%f, ", v)
	}
}

func fct2(matrix [][]float64) {
	matrix[2][0]= 12345.6
}

// Ok, lets implement sorter for 1D slices
// we could just use the sort built in functions
func sortFunc(tab []float64) ([]float64){
	return sort.Float64s(tab)


}

// Now we write the function to transpose a matrix
//func transpose(tab [][] float64) ([][]float64){

//}

func main() {
// array := [][]float64{{7.1, 2.3, 1.1},
// {4.3, 5.6, 6.8},
// {2.3, 2.7, 3.5},
// {4.5, 8.1, 6.6}}
array := [][]float64{{1.1, 7.3, 3.2, 0.3, 3.1},
					{4.3, 5.6, 1.8, 5.3, 3.1},
					{1.3, 2.7, 3.5, 9.3, 1.1},
					{7.5, 5.1, 0.6, 2.3, 3.9}}

//fct2(array)
//fct(array[2][:])
//fmt.Println(array[2][:])
sortFunc(array[2][:])
fmt.Println(array[2][:])
}