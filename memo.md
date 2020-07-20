# 技術的なめも

## main 関数の中でのエラーハンドリング
- Web サーバーなしの場合
→Fprintln + os.Exit
- ありの場合
→log.Fatal

## DB について
- main 関数の中で Open
db, err := sql.Open()
- 各関数の中で
Query, Exec
- rows
rows.Next, rows.Scan, rows.Err, defer rows.Close()


## エラーハンドリングについて
- main 関数の中
→log.Fatal
- 各関数の中
→あくまでも return にエラーを入れる処理

## handler について
これはカメラロールを参照
ServeHTTP との関係も

## 実態への関連付け
Scan 

## http.Error について
- code に http.StatusInternalServerError
- http.Error(w, "ほげ", code)
