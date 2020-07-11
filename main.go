package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.Category)

	fmt.Print("値段>")
	fmt.Scan(&item.Price)

	// ファイルに書き出し
	line := fmt.Sprintf("%s %d\n", item.Category, item.Price)

	if _, err := file.WriteString(line); err != nil {
		return err
	}
	return nil
}

// 一覧の表示
func showItems() error {
	file, err := os.Open("accountbook.txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// 1行ずつ取り出して、成形して、出力
		// ex. 食材 100 → 食材:100円
		line := fmt.Sprint(scanner.Text())
		splited := strings.Split(line, " ")
		if len(splited) != 2 {
			return errors.New("パースに失敗しました")
		}

		catefory := splited[0]
		strprice := splited[1]

		price, err := strconv.Atoi(strprice)
		if err != nil {
			return err
		}

		fmt.Printf("%s:%d円\n", catefory, price)
	}

	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}
