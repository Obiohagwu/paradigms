package main

import (
	"fmt"
	//"math"
)

type Point struct {
	x float64
	y float64
}

// Equation for coordinates of midpoint of a line:
// midpoint = ((x1+x2/2), (y1+y2/2))

// The questoin qants us to find the midpoint between every pair of points

// We are going have ot use a bit of math from the school of combinatorics
// Loop through lenght of sequence then use xCr to find combinations and midpoints
func length(p1 Point, p2 Point) float64{
	return ((p1.x+p2.x)/2) + ((p1.y+p2.y)/2)/2

}

func midpoint(p1 Point, p2 Point) Point{

	return Point{(p1.x+p2.x)/2, (p1.y+p2.y)/2}
}






func main() {
	//points := []Point{{8., 1.},{3., 2.},{7., 4.},{6., 3.}}
	points1 := []Point{{2., 4.},{6., 8.},{7., 4.},{6., 3.}}
	midpoints := make([]Point, len(points1))

	//fmt.Printf("point= %v\n", points[1:][0])
	for i:=0; i<len(points1)-1; i++ {
		mid :=midpoint(points1[i], points1[i+1])
		fmt.Println("Points: ", points1[i], points1[i+1])
		fmt.Printf("MidPoint= %v\n", mid)
		midpoints = append(midpoints, mid)
		
		
	}
	fmt.Println(midpoints)
	
}
