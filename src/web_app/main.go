package main

import (
	"net/http"
	"html/template"
	//"./data"
)

func index(w http.ResponseWriter, r *http.Request){
	files := []string{
		"templates/layout.html",
		"templates/main-description.html",
		"templates/navbar-collapse.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.Execute(w, nil)
}

func main(){
	//マルチプレクサの作成
	mux := http.NewServeMux()
	//HTTPリクエストに対して、第一引数のrootを起点とするファイルシステムのコンテンツを返すハンドラを返す
	files := http.FileServer(http.Dir("public/"))
	//URLのパスからプレフィックスを削除
	mux.Handle("/public/", http.StripPrefix("/public/", files))
	//ルートURLにハンドラ関数にリダイレクト
	mux.HandleFunc("/", index)

	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}