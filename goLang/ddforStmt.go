package main

import "fmt"

func main() {

//simplest for statement
    i := 1
    for i <= 3 {
        fmt.Println("First loop: ", i)
        i = i + 1
    }
        fmt.Println("out first loop: ",  i)

//traditional for statement but no parenthesis
    for j := 7; j <= 9; j++ {
        fmt.Println("Second loop: ",  j)
    }
 //fmt.Println("Second loop: ",  j)

//declare an array
	strategies := [5]int{1, 2, 3, 8, 5}//
        
//for k:= strategies{
  // fmt.Println("test", k);
//}
//k becomes indexes, not values in strategies
	for k := range strategies {
		fmt.Println("checkLoop ", strategies[k], k*2);

		strategies[k] =  k*2;
	}			
       for i := range strategies{
		fmt.Println("Third Loop ", + i);
	}
       for i := range strategies{
	    fmt.Println("fourth Loop ", + strategies[i]);
       }

    for i, c := range strategies {
        fmt.Println(i, c)
    }
}
