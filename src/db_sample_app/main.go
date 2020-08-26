package main


import (
	"fmt"
	"database/sql"
	// postgres ドライバ
	_ "github.com/lib/pq"
)

type TestUser struct{
	UserID int
	Password string
}

func main(){

	// Db: データベースに接続する為のハンドラ
	var Db *sql.DB
	Db, err := sql.Open("postgres", "host=postgres user=user password=secret1234 dbname=app_db sslmode=disable")
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	sql := "SELECT user_id, user_password FROM TEST_USER WHERE user_id=$1;"

	pstatement, err := Db.Prepare(sql)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}

	//検索パラメータ（ユーザーID）
	queryID := 1
	//検索結果用変数
	var test_user TestUser

	err = pstatement.QueryRow(queryID).Scan(&test_user.UserID, &test_user.Password)
	if err != nil {
		fmt.Printf("ERROR: %d", err)
	}

	fmt.Printf("Hello Worlld, %s!", test_user.UserID)

}