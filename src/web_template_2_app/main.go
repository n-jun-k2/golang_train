package main

import(
	"net/http"
	"html/template"
	"log"
	"time"
)
func index(w http.ResponseWriter, r *http.Request){

	files := []string{
		"templates/index.html",
		"templates/_footer.html",
		"templates/_header.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalf("template error: %v", err)
	}

	if err := t.Execute(w, struct{
		Title string
		Message string
		Time time.Time
	}{
		Title: "テストページ",
		Message: "こんにちは!",
		Time: time.Now(),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}

}

func main() {
	// マルチプレクサの生成を行う。
	mux := http.NewServeMux()
	// マルチプレクサの静的なファイル返送を行う。
	files := http.FileServer(http.Dir("/public"))
	// リクエストのURLパスからプレフィックスを削除する /satic/で始まる全てのリクエストURLから文字列/static/を取り去る。
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	/* ルートURLをハンドラ関数にリダイレクトする
		第一引数：URL
		第二引数：ハンドラ関数名
	*/
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}