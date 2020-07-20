package main

import (
	"html/template"
	"net/http"
)

// Handlers :http ハンドラーを集めた型
type Handlers struct {
	ab *AccountBook
}

// NewHandlers :Handlers を作成
func NewHandlers(ab *AccountBook) *Handlers {
	return &Handlers{ab: ab}
}

// TODO: 別ファイルに移植
// ListHandlerで仕様するテンプレート
var listTmpl = template.Must(template.New("list").Parse(`<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8"/>
		<title>家計簿</title>
	</head>
	<body>
		<h1>家計簿</h1>
		<h2>最新{{len .}}件</h2>
		{{- if . -}}
		<table border="1">
			<tr><th>品目</th><th>値段</th></tr>
			{{- range .}}
			<tr><td>{{.Category}}</td><td>{{.Price}}円</td></tr>
			{{- end}}
		</table>
		{{- else}}
			データがありません
		{{- end}}
	</body>
</html>
`))

// ListHandler :最新の入力データを表示するハンドラ
func (hs *Handlers) ListHander(w http.ResponseWriter, r *http.Request) {
	items, err := hs.ab.GetItems(10)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := listTmpl.Execute(w, items); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
