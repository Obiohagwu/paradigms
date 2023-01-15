package main

import (
	"fmt"
	"math"
)

func ex1(n float64) (float64, float64) {
	return math.Ceil(n), math.Floor(n)
}


func main(){
	fmt.Println(ex1(2.3))

}

