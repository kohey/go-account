package main

import "fmt"

// Item で品目と値段を扱う
type Item struct {
	Category string
	Price    int
}

func main() {
	item := inputItem()
	fmt.Println("===========")
	fmt.Printf("%s に %d 円使いました", item.Category, item.Price)
	fmt.Println("===========")
}

func inputItem() Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	return item
}
