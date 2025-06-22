package main
import "fmt"

func main(){
	age:= 5
	if age>= 18 {
		fmt.Println("adult")
	} else if age>= 12{
		fmt.Println("teeneger")
	}else{
		fmt.Println("kid")
	}

var role ="admin"
var hasper = false

if role=="admin" && hasper==true{
	fmt.Println("yes")
}else{
	fmt.Println("no")
}
//initailization and condition in same line
//we can declare varible inside if construct
if year:=20; year>=18{
	fmt.Println("adult",year)
}else{
	fmt.Println("not adult", year)
}
//go does not have ternary, will have to use normal if else
}