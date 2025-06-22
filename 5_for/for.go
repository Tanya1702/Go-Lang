package main

import "fmt"

// for -> only construct in go for looping
func main(){


 //while loop fashion in for loop

 i:= 1
 for i<=3{
	fmt.Println(i)
	i= i+1 }

// for{
// 	fmt.Println(1)
// }

//classic for loop

for j:=0;j<=3;j++{
	//break
	if j==2 {
		continue
	}
	fmt.Println(j)
}

for k:=range 3 {
	//new feature that is range
	fmt.Println(k)
}
 }
