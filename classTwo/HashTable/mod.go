package main

import "fmt"


func main()  {
	a := 18 //1000
	b := 5 //0011    
	fmt.Println(a%b)
	fmt.Println(a&(b-1)) // a & （b-1） 相当于 a 最后散列到 b 范围内，不等于a%b
} 