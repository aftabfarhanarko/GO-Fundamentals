package main

import "fmt"

func main() {
	fmt.Println("Hello GO")
	//
	// Comparison
	fmt.Println(10 == 10)
	fmt.Println(10 == 29)
	fmt.Println(20 != 30)
	fmt.Println(20 > 99)
	fmt.Println(30 < 40)
	fmt.Println(89 >= 89)
	fmt.Println(45 <= 45)

	// Logical Oprerators
	fmt.Println(45 > 42 && 89 < 90)
	fmt.Println(56 > 34 || 78 > 88)

	i := 56

	i++
	fmt.Println(i)

	i--
	fmt.Println(i)

}
