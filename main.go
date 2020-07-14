package main

import (
	"fmt"
)

// 入力によって処理を分岐
func main() {

}

// item に 入力値を紐づけて、返す
func inputItem(item Item) *Item {
	fmt.Print("品目>")
	fmt.Scan(&item.Category)
	fmt.Print("価格>")
	fmt.Scan(&item.Price)

	return &item
}

// Items の一覧を、「〇:〇円」という形で返すだけ
func showItems(items []*Item) {
	for _, item := range items {
		fmt.Printf("%s: %v", item.Category, item.Price)
	}
}
