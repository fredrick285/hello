package main

import (
	"fmt"
)

func SayHello(names []string) {
	for _, name := range names {
		fmt.Println("Hi, " + name)
	}
}
func main() {
	SayHello([]string{"Alice", "Bob", "Charlie"})
	// fmt.Println("Hello, Go on WSL!")//
	//var nameOne string = "Emmy"
	//var numOne = 10
	//var nameTwo string
	//nameFour := "peaches"
	//fmt.Println(nameOne, numOne, nameTwo, nameFour)

	//var numIne int8 = 25
	//var numTwo int8 = -128

	//var numThree int8 = 129
	///fmt.Println(numIne, numTwo, numThree)
	//age := 30
	//name := "Eddy"
	//fmt.Printf("My name is %v and my age is %v", name, age)
	//	var Ages = [4]int{20, 30, 40, 50}
	////	Ages[0] = 25
	//	var NewAge = []int{60, 100, 70, 900}
	//	var NewAges = append(NewAge, 90)
	//	fmt.Println(Ages, NewAges)
	//	rangeOne := NewAges[0:3]
	//	fmt.Println(rangeOne)
	//x := 13
	//for x < 31 {
	//	fmt.Println("Value of x:", x)
	//	x++
	//}
	//var Names = []string{"Eddy", "Emmy", "Peaches", "Mango"}
	//for i, name := range Names {
	//	fmt.Println(i, name)
	//	i++
	//}
	//for i := 1; i < len(Names); i++ {
	//	fmt.Println(i, Names[i])
}
