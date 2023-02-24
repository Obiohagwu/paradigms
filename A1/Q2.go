// Michael Ohagwu   
// 300074813

package main
import (
	"fmt"
	"sync"
	//"sort"
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

// Bubble sort is easy and fairly quick to implement
func sortFunc(tab []float64){
	n := len(tab)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if tab[j] > tab[j+1] {
                tab[j], tab[j+1] = tab[j+1], tab[j] // swap the elements
            }
        }
    }

}

// Now we write the function to transpose a matrix
// A transpose is basically a swap between columns and rows

func transpose(tab [][] float64) {
    transposed := make([][]float64, len(tab[0]) )
    for i := range transposed {
        transposed[i] = make([]float64, len(tab))
    }
    for i := 0; i < len(tab); i++ {
        for j := 0; j < len(tab[0]) ; j++ {
            transposed[j][i] = tab[i][j]
        }
    }

}

func sortRows(tab [][]float64){
	var waitgroup sync.WaitGroup
	for i := range tab {
		waitgroup.Add(1)
		go func (row []float64)  {
			defer waitgroup.Done()
			sortFunc(row)
		}(tab[i])

	}
	waitgroup.Wait()
}


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
//sortFunc(array[2][:])
fmt.Println("Input:", array)
sortRows(array)
transpose(array)
sortRows(array)
transpose(array)
fmt.Println("Output:", array)


//fmt.Println(array[2][:])
}