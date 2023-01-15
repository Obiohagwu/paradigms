package main

import "fmt"

// EX3 STRUCTS & POINTERS
type Person struct {
	lastName  string
	firstName string
	iD        int
}

func inPerson(p *Person, lastId int) (nextId int, err error) {
	// Input parameters p of type Person and lastId of type int
	nextId = lastId

	//To enter lastName from console user input
	fmt.Print("Last Name: ")
	_, err = fmt.Scanf("%s", &p.lastName) // This says to scan input from console and send to lastName of person p
	if err != nil {
		return
	}

	//To enter first name. We repeat what we did for lastName above
	fmt.Print(("First Name: "))
	_, err = fmt.Scanf("%s", *&p.firstName)
	if err != nil {
		return
	}

	//Then to generate the ID per person p
	nextId += 1
	p.iD = nextId
	return
}

func printPerson(p Person) {
	fmt.Printf("%10.d\t%s\t%s\n", p.iD, p.firstName, p.lastName)

}

func main(){
	nextId := 101
	for {
		var (
			p   Person
			err error
		)
		nextId, err = inPerson(&p, nextId)
		if err != nil {
			fmt.Println("Invalid entry ... exiting")
			break
		}
		printPerson(p)
	}

}