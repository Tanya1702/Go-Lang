package main

import "fmt"

const company = "accenture"
func main(){
// keyword for constant is const
	const name= "tanya"
//name="gupta" not possible
	fmt.Println(name)

	const age = 25
	car:= "tata"
	fmt.Println(age)
	fmt.Println(company)
	fmt.Println(car)

	// multiple constants

	const(
		port = 500
		host = "local"
	)

	fmt.Println(port,host)
}