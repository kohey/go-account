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
		fmt.Fprintln(os.Stderr, "dbオープンエラー:", err)
		os.Exit(1)
	}
	ab := NewAccountBook(db)

	// 存在しなかったらテーブルを作成
	if err := ab.CreateTable(); err != nil {
		fmt.Fprintln(os.Stderr, "table 作成エラー", err)
	}

LOOP:
	for {
		var mode int
		fmt.Println("[1]入力 [2]最新10件 [3]集計 [4]終了")
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
				break LOOP
			}

			showItems(items)

		case 3:
			summeries, err := ab.GetSummeries()
			if err != nil {
				fmt.Fprintln(os.Stderr, "集計エラー", err)
				break LOOP
			}
			showSummery(summeries)
		// 終わりたい場合
		case 4:
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

// Items の一覧を返す
func showItems(items []*Item) {
	for _, item := range items {
		fmt.Printf("[%04d]%s:%d円", item.ID, item.Category, item.Price)
	}
}

// 集計結果を出力する
func showSummery(summeries []*Summery) {
	fmt.Printf("品目\t個数\t合計\t平均\n")

	for _, s := range summeries {
		fmt.Printf("%s\t%d\t%d円\t%.2f円\n", s.Category, s.Count, s.Sum, s.Avg())
	}
}
