package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name   string
	age    int
	height float64
}

func main() {
	// longhand
	var x int = 12
	// Shorthand inference of type
	y := 30
	sum := x + y

	fmt.Println(sum)

	//Longhand
	var myArray [5]int
	myArray[2] = 13
	fmt.Println(myArray)

	//shorthand
	myArray2 := [5]int{1, 34, 76, 233, 4}
	fmt.Println(myArray2)

	//Slice instead of array
	mySlice := []int{23, 65, 67}
	mySlice = append(mySlice, 654)
	fmt.Println(mySlice)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["circle"] = 0
	vertices["hexagon"] = 6

	fmt.Println(vertices)
	fmt.Println("A triangle has", vertices["triangle"], "vertices")

	delete(vertices, "circle")
	fmt.Println(vertices)

	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	for key, value := range vertices {
		fmt.Println("key:", key, "value:", value)
	}

	result := mySum(3, 23)
	fmt.Print("The sum of 3 + 23 = ", result)

	badNum := -100.0
	goodNum := 1000.0
	mySqrt1, err1 := sqrt(badNum)
	mySqrt2, err2 := sqrt(goodNum)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Square root of", badNum, "is", mySqrt1)
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Square root of", goodNum, "is", mySqrt2)
	}

	p := person{name: "Wade", age: 40, height: 69.5}
	fmt.Println(p)

	myInc := 34
	// Pass by memory reference
	inc(&myInc)
	fmt.Println(myInc)
}

func mySum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers.")
	}

	return math.Sqrt(x), nil
}

func inc(x *int) {
	//Dereference
	*x++
}
