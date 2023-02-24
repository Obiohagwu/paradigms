// Michael Ohagwu   
// 300074813

package main
import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Equation for coordinates of midpoint of a line:
// midpoint = ((x1+x2/2), (y1+y2/2))

// length is simply given by pythagorean equation
// The questoin wants us to find the midpoint between every pair of points

// We are going have ot use a bit of math from the school of combinatorics
// Loop through lenght of sequence then use xCr to find combinations and midpoints

func MidPoint(p1 Point, p2 Point) (Point, float64) {
    coord_x := (p1.x + p2.x) / 2.0
    coord_y := (p1.y + p2.y) / 2.0
	midPoint := Point{coord_x, coord_y}
    length := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
    //fmt.Printf("Points: (%v,%v) (%v,%v)\n", p1.x, p1.y, p2.x, p2.y)
    //fmt.Printf("MidPoint= (%.2f, %.2f)\n", coord_x, coord_y)
    //fmt.Printf("Length= %.2f\n", length)
	return midPoint, length
}





func main() {
	//points := []Point{{8., 1.},{3., 2.},{7., 4.},{6., 3.}}
	points := []Point{{2., 4.},{6., 8.},{7., 4.},{6., 3.}}
    //ch := make(chan bool, len(points)*(len(points)-1)/2)
	ch := make(chan string)

	// The questoin wants us to find the midpoint between every pair of points
	// We are going have ot use a bit of math from the school of combinatorics
	// Loop through lenght of sequence then use xCr to find combinations and midpoints

    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            go func(p1 Point, p2 Point) {
               midPoint, length := MidPoint(p1, p2)
			   result := fmt.Sprintf("Points: (%.0f, %.0f) (%.0f, %.0f)\nMidPoint= (%.1f, %.1f)\nLength= %.2f\n", p1.x, p1.y, p2.x, p2.y, midPoint.x, midPoint.y, length)
			   ch <- result
			 
            }(points[i], points[j])
        }
    }

    for i := 0; i < len(points)*(len(points)-1)/2; i++ {
        fmt.Print(<-ch)
    }
	
}
