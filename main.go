package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/tenntenn/sqlite"
)

// 入力によって処理を分岐
func main() {
	db, err := sql.Open(sqlite.DriverName, "accountbook.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ab := NewAccountBook(db)

	// 存在しなければ、テーブルを新規作成
	if err := ab.CreateTable(); err != nil {
		log.Fatal(err)
	}

	// handlers の新規作成
	hs := NewHandlers(ab)

	// handlers の登録
	http.HandleFunc("/", hs.ListHander)
	http.HandleFunc("/save", hs.SaveHandler)

	// サーバーの開始
	fmt.Print(":8080ポートで WebServer を起動中・・・")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
