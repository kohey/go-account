package main

import "fmt"

func main() {
	var category string
	var price int

	fmt.Print("品目>")
	fmt.Scan(&category)

	fmt.Print("値段>")
	fmt.Scan(&price)

	fmt.Println("============")

	fmt.Printf("%s に %d 円使いました", category, price)

	fmt.Println("============")
}
