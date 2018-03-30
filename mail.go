package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/beeker1121/mailchimp-go/lists/members"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	connStr := "user=postgres host=localhost dbname=teratail sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(db)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("データベースの接続に失敗しました。2: %v", err)
	}
	rows, err := db.Query("SELECT user_id, setting FROM user_mail_settings")
	id := ""
	settings := ""
	setting := []string{}
	for rows.Next() {
		err = rows.Scan(&id, &settings)
		checkError(err)
		fmt.Println(id, settings)

		json.Unmarshal(settings, &setting)
	}
}

func main1() {
	member, err := members.GetMember("123456", "123", nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(member)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}
