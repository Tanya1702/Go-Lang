package main
import "fmt"

// array is numbered sequence of specific length
//array have fixed size , memory optimization,constant time access
func main(){
	//zeroed values int->0 string->" " bool->false

	//array declare
	var nums[4]int
	
	//add element
	nums[0] =10

	//display element
	fmt.Println(nums[0])
	fmt.Println(nums)

    //to get array length
	fmt.Println(len(nums))

	var bools[5]bool
	bools[2]= true
	fmt.Println(bools)

	var name[4]string
	name[0]="tanya"
	name[1]="gupta"
	name[2]="aa"
	name[3]="bb"
	fmt.Println(name)

	num :=[3]int{5,4,3}
	fmt.Println(num)

	//2d array

	arr :=[2][2]int{{1,2},{3,4}}
	fmt.Println(arr)


}