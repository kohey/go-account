package main

import (
	"database/sql"
)

// Item deals Category and Price
type Item struct {
	ID       int
	Category string
	Price    int
}

// AccountBook :家計簿の処理を行う型
type AccountBook struct {
	db *sql.DB
}

// NewAccountBook :新しい AccountBook を作成する
func NewAccountBook(db *sql.DB) *AccountBook {
	return &AccountBook{db: db}
}

// CreateTable :存在しなければ、新しい itemsテーブルを作成
func (ab *AccountBook) CreateTable() error {
	const sql = `CREATE TABLE IF NOT EXISTS items(
								id INTEGER PRIMARY KEY,
								category STRING NOT NULL,
								price INTEGER NOT NULL);
							`
	// 実行する sql の準備
	_, err := ab.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

// AddItem :items テーブルに新しい Item を追加する
func (ab *AccountBook) AddItem(item *Item) error {
	const sql = `INSERT INTO items(category, price)	VALUES(?,?)`

	_, err := ab.db.Exec(sql, item.Category, item.Price)
	if err != nil {
		return err
	}

	return nil
}

// GetItems :DBに最近追加したものを、item に紐付けて、 limit 分返す
func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {
	const sql = `SELECT * FROM items ORDER BY id DESC`
	rows, err := ab.db.Query(sql)
	if err != nil {
		return nil, err
	}

	var items []*Item
	defer rows.Close()
	// 取得したレコードそれぞれについて、item に割り当て
	// items に追加
	for rows.Next() {
		var item Item
		err := rows.Scan(
			&item.ID,
			&item.Category,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	// rows.Err: rows.Next のループ中に起きた様々なエラーを取得
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// GetSummeries :品目ごとに金額を集計して返す
func (ab *AccountBook) GetSummeries() ([]*Summery, error) {
	const sql = `
	SELECT Category, Count(1) as count, Sum(price) as sum,
	FROM items
	GROUY BY Category
	`
	rows, err := ab.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var summeries []*Summery

	for rows.Next() {
		var summery Summery
		rows.Scan(
			&summery.Category,
			&summery.Count,
			&summery.Sum,
		)
		summeries = append(summeries, &summery)
	}

	// ループの終了のハンドリング
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return summeries, nil
}

// Summery :品目ごとにまとめたデータ
type Summery struct {
	Category string
	Count    int
	Sum      int
}

// Avg :全品目の平均
func (s *Summery) Avg() float64 {
	if s.Count == 0 {
		return 0
	}

	return float64(s.Sum) / float64(s.Count)
}
