package main

import "fmt"


func statusPrint(state int8) {
	switch state {
	case 0:
		fmt.Printf("State is %s (%d)\n", "Idle", 0)
	case 1:
		fmt.Printf("State is %s (%d)\n", "Start", 1)
	case 2:
		fmt.Printf("State is %s (%d)\n", "Forward", 2)
	case 3:
		fmt.Printf("State is %s (%d)\n", "Fast", 3)
	case -1:
		fmt.Printf("State is %s (%d)\n", "Reverse", -1)
	default:
		fmt.Printf("Invalid state: %d\n", state)
	}
	return
}

func main() {
	var i int8
	for i = -1; i < 5; i++ {
		statusPrint(i)
	}
}