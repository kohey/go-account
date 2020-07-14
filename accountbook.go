package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Item deals Category and Price
type Item struct {
	Category string
	Price    int
}

// AccountBook :家計簿の処理を行う型
type AccountBook struct {
	fileName string
}

// NewAccountBook :新しい AccountBook を作成する
func NewAccountBook(fileName string) *AccountBook {
	return &AccountBook{fileName: fileName}
}

// AddItem :ファイルに新しい Item を追加する
func (ab *AccountBook) AddItem(item *Item) error {
	file, err := os.OpenFile(ab.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 「品目 価格」の順番に記録
	if _, err := fmt.Fprintln(file, item.Category, item.Price); err != nil {
		return err
	}

	// ファイルを閉じる
	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

// GetItems :ファイルに記載してある、最近追加したものを、item に紐付けて、 limit 分返す
func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {
	file, err := os.Open(ab.fileName)
	if err != nil {
		return nil, err
	}

	// ファイルをスキャン
	scanner := bufio.NewScanner(file)

	var items []*Item
	// 1行ずつ読みこむ
	for scanner.Scan() {
		var item Item

		if err := ab.parseLine(scanner.Text(), &item); err != nil {
			return nil, err
		}

		// この時点で item に値が紐づけられているので
		items = append(items, &item)
	}

	// limit より少なかったら、全部返す
	if len(items) < limit {
		return items, nil
	}

	// 後方から順番に limit 分取り出し
	stIndex := len(items) - limit
	enIndex := len(items)

	return items[stIndex:enIndex], nil
}

// 1行をパースする
// パース = 1行を item(実体) に割り当てる
func (ab *AccountBook) parseLine(line string, item *Item) error {
	splited := strings.Split(line, " ")
	// 分割がおかしかったら、終了
	if len(splited) != 2 {
		return errors.New("parse に失敗しました")
	}

	category := splited[0]

	price, err := strconv.Atoi(splited[1])
	if err != nil {
		return err
	}

	item.Category = category
	item.Price = price

	return nil
}
