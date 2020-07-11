package main

import "fmt"

// Item で品目と値段を扱う
type Item struct {
	Category string
	Price    int
}

func main() {
	var n int
	fmt.Print("品目数>")
	fmt.Scan(&n)

	items := make([]Item, 0, n)

	for i := 0; i < cap(items); i++ {
		items = inputItem(items)
	}

	showItems(items)
}

// 入力の受付
func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	items = append(items, item)
	return items
}

// 一覧の表示
func showItems(items []Item) {
	fmt.Println("===========")
	// 全ての item を表示
	for i := 0; i < len(items); i++ {
		fmt.Printf("%s: %d 円", items[i].Category, items[i].Price)
	}
	fmt.Println("===========")
}
