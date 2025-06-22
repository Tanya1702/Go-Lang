package main
import "fmt"
import "time"


func main (){
	//simple switch
	i:=20

	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("other")	
	}

	// multiple condition switch
	// use standard library package that is time

	switch time.Now().Weekday(){

		case time.Saturday, time.Sunday:
			fmt.Println("weekend")
		default:
			fmt.Println("workday")
	}

	//type switch
	//interface{}-- means any data type
	  whoAmI := func(i interface{}){
		switch t := i.(type){   //i.(type) will return the data type of varible
	      case int:
		   fmt.Println("integer")
    	  case string:
		   fmt.Println("string")
	      case bool:
		   fmt.Println("boolean")
	      default:
		   fmt.Println("other",t)
	  }

	  }

whoAmI("TANYA")
}
