package main

import (
	"database/sql"
	"log"
	"crypto/sha256"
	_ "github.com/go-sql-driver/mysql"
)

func addText(contents string, source int) {
	byteArray := sha256.Sum256([]byte(contents))

	stmt, err := db.Prepare("INSERT INTO text (contents, source, hash) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(contents, source, byteArray[:])
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func addWordNatural(word string, definitionText string, source string) {
}

func addWord(word string, definitionTextHash []byte) {
	stmt, err := db.Prepare(`INSERT INTO word (word, definition)
SELECT 
    id AS definition FROM text
    word FROM ?
WHERE hash = ?`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(word, definitionTextHash)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func addWordDacoda(word string, definitionTextId int) {
	stmt, err := db.Prepare("INSERT INTO word (word, definition) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(word, definitionTextId)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func addNewChildWordByIds(word string, definition string, source string, user string, parent_text []byte) {
	addWord()

}

func addChildWordByIds(user int, parent_text int, word int) {

}

func listSources() {

	var (
		id int
		name string
	)

	rows, err := db.Query("SELECT id, name FROM source");

	defer rows.Close();

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

var db, err = sql.Open("mysql", "root:dacodastrackoda@tcp(127.0.0.1:3306)/orihime")

func main() {
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	addWord("Fuck you", []byte{1, 1, 1, 1})
	// addText("Who art thou", 1);
}
