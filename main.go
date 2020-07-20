package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/tenntenn/sqlite"
)

// 入力によって処理を分岐
func main() {
	db, err := sql.Open(sqlite.DriverName, "accountbook.db")
	if err != nil {
		// 標準出力に書き出して終了
		fmt.Fprintln(os.Stderr, "エラー:", err)
		os.Exit(1)
	}
	ab := NewAccountBook(db)

LOOP:
	for {
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]終了")
		fmt.Print(">")
		fmt.Scan(&mode)

		switch mode {
		// 新規入力の場合
		case 1:
			var n int
			fmt.Println("何件入力しますか？>")
			fmt.Print(">")
			fmt.Scan(&n)

			for i := 0; i < n; i++ {
				var item Item
				if err := ab.AddItem(inputItem(item)); err != nil {
					fmt.Fprintln(os.Stderr, "エラー:", err)
					break LOOP
				}
			}
		// 最新 item の取得の場合
		case 2:
			var limit int
			fmt.Println("何件取得しますか？>")
			fmt.Print(">")
			fmt.Scan(&limit)

			// limit 分の item を取得
			items, err := ab.GetItems(limit)
			if err != nil {
				fmt.Fprintln(os.Stderr, "エラー", err)
			}

			showItems(items)
			break LOOP
		// 終わりたい場合
		case 3:
			fmt.Println("終了します")
			return
		}
	}
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
