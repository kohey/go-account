package main

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
}

// GetItems :最近追加したものを limit 分返す
func (ab *AccountBook) GetItems(limit int) ([]*Item, error) {
}

// 1行ずつパースする
func (ab *AccountBook) parseLines(line string, item *Item) error {

}
