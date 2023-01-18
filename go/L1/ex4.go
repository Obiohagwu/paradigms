package main

import "fmt"

// Define the structs

type Course struct {
	Professor string
	NStudents int64
	Avg float64
}




func main(){
	// Create dynamic map m
	m := make(map[string]Course)


	// Add the courses CSI2120 and CSI2110 to the map
	m["CSI2120"] = Course{"Lang", 186, 79.5000}
	m["CSI2110"] = Course{"Moura", 211, 81.000}

	for k, v := range m{
		fmt.Printf("Course Code: %s\n", k)
		fmt.Printf("Number of students: %d\n", v.NStudents)
		fmt.Printf("Professor: %s\n", v.Professor)
		fmt.Printf("Average: %f\n\n", v.Avg)
	}
}