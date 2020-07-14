package main

import (
	"fmt"
	"log"
	"os"
)

// Item で品目と値段を扱う
type Item struct {
	Category string
	Price    int
}

func main() {
	file, err := os.Create("accountbook.txt")
	if err != nil {
		log.Fatal(err)
	}

	var n int
	fmt.Print("何件入力しますか？>")
	fmt.Scan(&n)

	// 指定された回数、入力を受け付ける
	for i := 0; i < n; i++ {
		if err := inputItem(file); err != nil {
			// エラーしたら終了
			log.Fatal(err)
		}
	}

	// ファイルを閉じる
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	// ファイルに書いてある分は、全て出力する
	if err := showItems(); err != nil {
		log.Fatal(err)
	}
}

// 入力を受け付けて、ファイルに保存する
func inputItem(file *os.File) error {

}

// Items の一覧を、「〇:〇円」という形で返すだけ
func showItems(items []*Item) {
	for _, item := range items {
		fmt.Printf("%s: %v", item.Category, item.Price)
	}
}
