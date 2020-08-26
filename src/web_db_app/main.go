package main

import (
	"net/http"
	"html/template"
	"log"
	"time"

	"fmt"
	"database/sql"
	// postgres ドライバ
	_ "github.com/lib/pq"
)


var Db *sql.DB

func printLog(err error){
	if err != nil{
		fmt.Printf("Error: %s", err)
	}
}

func generate_web(w http.ResponseWriter, files []string, str interface {}){
	templates := template.Must(template.ParseFiles(files...))
	if err := templates.Execute(w, str); err != nil {
		log.Printf("failed to execute template: 5v", err)
	}
}

func query_send(query string, param ... interface {}) []string {

	Db, err := sql.Open("postgres", "host=postgres user=user password=secret1234 dbname=app_db sslmode=disable")
	printLog(err)

	pstatement, err := Db.Prepare(query)
	printLog(err)

	prow, err := pstatement.Query(param...)
	printLog(err)
	defer prow.Close()

	cols, err := prow.Columns()
	printLog(err)

	//出力結果保持用
	result := []string{}
	//バイト配列をScanのインターフェースに渡す為のバッファポインタ
	dest := make([]interface{}, len(cols))
	//バイト配列でScanの結果を取得する為のバッファ
	raws := make([][]byte, len(cols))

	//destとrawを結びつける
	for i, _ := range raws {
		dest[i] = &raws[i]
	}

	for prow.Next(){
		serr := prow.Scan(dest...)
		printLog(serr)

		//byte配列をstringに変換
		for _, raw := range raws {
			if raw != nil {
				result = append(result, string(raw))
			}
		}
	}
	return result
}

func index(w http.ResponseWriter, r *http.Request) {

	itemlist := query_send("SELECT user_id, user_password FROM TEST_USER;")
	fmt.Printf("%#v\n", itemlist)

	generate_web(w, []string{
		"templates/layout.html",
		"templates/copyright.html",
	}, struct{
			Title string
			Message string
			Time time.Time
			Itemlist []string
		}{
			Title: "テストページ",
			Message: "こんにちは",
			Time: time.Now(),
			Itemlist: itemlist,
		})
}


func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}