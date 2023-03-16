package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var htmlStr string

func main() {
	fmt.Println("Server Started!")
	fmt.Println("Access to http://localhost:8080 and enjoy your project!")

	// 外部ファイルを読み込むときはReadFile関数
	// golangは関数の戻り値を複数取得することが可能
	// 外部からデータを取得する場合はこの書き方が一般的
	data, err := os.ReadFile("index.html")
	// データが取得できない場合はエラーログを返す
	if err != nil {
		log.Fatal(err)
	}
	htmlStr = string(data)

	// 第一引数：パス、第二引数：実行する関数などの処理
	http.HandleFunc("/", home)
	// ポート番号を指定してサーバーを起動	
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// ブラウザがfavicon.icoのリクエストを自動で行ってしまうため余計にhandler関数が走ってしまうことを避ける
	if r.URL.Path == "/favicon.ico" {
        http.ServeFile(w, r, "favicon.ico")
        return
    }
	fmt.Fprintln(w, htmlStr)
}