package main
import (
	"fmt"
	"slices"
)
//slice are dynamic arrays , most used construct in go , +useful methods
//during resize it will double its current capasity

func main(){

	//declare slice , uninitialized slice is nil
	 var nums[]int
	fmt.Println(nums)
	fmt.Println(nums==nil)
	fmt.Println(len(nums))

	//to initialize value: use make func
	//x{name of array varible} := make{allocate space in memory}([]int{data type},length,capacity)
	nums = make([]int, 0,5)
	// 2nd paramenter -slick will be of that length and valued element will be added their
	//cap: capacity-> max number of elements that can fit can be used in 3rd parameter as inital capasity
	fmt.Println(cap(nums))
	fmt.Println(nums)
	fmt.Println(nums==nil)

	//function to add element is append. 
	nums = append(nums,1)
	nums = append(nums,2)
	nums = append(nums,3)
	nums = append(nums,4)
	nums = append(nums,5)
	
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	//1:25:36

	//other way to make slices i.e, empty slice
	num :=[]int{}
	
	fmt.Println(num)
	fmt.Println(cap(num))
	fmt.Println(len(num))

	// to add number using index , if we add element in slice with length 0 then it will show error i.e in   numb = make([]int, 0,5)
	var numb[]int
	numb = make([]int, 2,5)
	numb[0]=3
	numb[1]=10
	fmt.Println(numb)

	fmt.Println(cap(numb))
	fmt.Println(len(numb))

	//copy function to copy a slice
	var cnums = make([]int,0,5)
	cnums = append(cnums,2)
	var cnums2 = make([]int , len(cnums))
	
	//copy function copy(destination,source)
	fmt.Println(cnums,cnums2)
	copy(cnums2,cnums)
	fmt.Println(cnums,cnums2)

	//slice operator   :(print elements from first index excluding 2nd index between given in slice)
	var onums = []int{1,2,3,4,5,6}
	fmt.Println(onums[0:2])
	fmt.Println(onums[:1])
	fmt.Println(onums[1:])
	fmt.Println(onums[:])

	//slice inbuilt package

	var pnum1 = []int {1,2,3}
	var pnum2 = []int {1,2,3}
	var pnum3 = []int {1,2,4}

	//slices.Equal() : to compare if 2 slices are equal or not

	fmt.Println(slices.Equal(pnum1,pnum2))
	fmt.Println(slices.Equal(pnum1,pnum3))

// to make 2d slices
var dnums = [][] int{{1,2,3},{4,5,6}}
fmt.Println(dnums)

}