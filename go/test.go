package main

import "fmt"


func main(){
	prefix := "RRCSIL"
	arr := []string{"3", "2", "1", "2", "0", "4"}
	suffix := func(arr []string)(concat string){
		for _, s := range arr[1:]{
			concat+=s
		}
		return

	}(arr)
	fmt.Print("Welcome to", prefix[2:5], suffix[:-1])
	
	
}