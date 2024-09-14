package main

import "fmt"

func add(x int, y int) (summa int) {
	summa = x + y
	return
}

func swap(first string, second string) (string, string) {
	return second, first
}

func quotientRemainder(x int, y int) (int, int) {
	return x / y, x % y
}

func main() {

	// fmt.Println("Hello World!")

	// var number int = 12
	// fmt.Printf("Value is %v, type is %T\n", number, number)

	// var price = 12.34
	// fmt.Printf("Value is %v, type is %T\n", price, price)

	// var sentence string
	// sentence = "I am man"
	// fmt.Printf("Value is %v, type is %T\n", sentence, sentence)

	// flag := true
	// fmt.Printf("Value is %v, type is %T\n", flag, flag)

	// fmt.Scan(&number)
	// if number == 0 {
	// 	fmt.Println("The number is 0")
	// } else if number < 0 {
	// 	fmt.Println("The number is negative")
	// } else {
	// 	fmt.Println("The number is positive")
	// }

	// sum := 0
	// for i := 1; i < 11; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// fmt.Scan(&number)
	// switch number {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Not a weekday")
	// }
	
	fmt.Println(add(1, 2))
	fmt.Println(swap("Bekzat", "Tursun"))
	fmt.Println(quotientRemainder(5, 4))
	
	

}
