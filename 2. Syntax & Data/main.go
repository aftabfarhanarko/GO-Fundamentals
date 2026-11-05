package main

import (
	"fmt"
	"strconv"
)

// var stringTow string = "Hello World"
// var numberTow int = 2
// var foltTow float64 = 20.67
// var boolTow bool = false

// Decliear Constinc
const (
	AppName = "StockFlow"
	Version = 1.0
	MaxUser = 500
)

// How TO  Declier IOTA Count a use  
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Syntax & Data Type")

	i1 := 100;
	f1 := float64(i1);

	fmt.Println(f1)

	f2  := 22.34;
	i2 := int(f2);
	fmt.Println(i2)

	i3 :=  65;
	s1 := strconv.Itoa(i3);
	fmt.Println(s1)


	y2 := "YGFYU";
	i5, _ := strconv.Atoi(y2);
	fmt.Println(i5)


	// var stringOne string = "Hello"
	// var numberOne int = 1
	// var foltOne float64 = 19.67
	// var boolOne bool = true

	// var stringOne  = "Hello sfqw qwfqwf "
	// var numberOne  = 1
	// var foltOne  = 19.67
	// var boolOne  = true

	// Short From Use Just Functions in side

	stringOne := "dfsfsf"
	numberOne := 45
	foltOne := 34.56
	boolOne := true

	// Change variable values
	stringOne = "Hello Go Lang Variable"
	numberOne = 3333
	foltOne = 00.0
	boolOne = false

	// Multiple variable declarations
	a, b := 3, 5
	u, i := "Arko", "Go"
	e, w := 99.5, 22.4
	o, p := false, true

	fmt.Println(stringOne)
	fmt.Println(numberOne)
	fmt.Println(foltOne)
	fmt.Println(boolOne)

	fmt.Println(a, b)
	fmt.Println(u, i)
	fmt.Println(e, w)
	fmt.Println(o, p)

	fmt.Println(AppName)
	fmt.Println(Version)
	fmt.Println(MaxUser)

	
    fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
