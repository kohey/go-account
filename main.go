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
	http.HandleFunc("/summary", hs.SummeryHandler)

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
