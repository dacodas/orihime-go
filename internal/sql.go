package internal

import (
	"log"
	"database/sql"
	"crypto/sha256"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:dacodastrackoda@tcp(127.0.0.1:3306)/orihime")

func checkDatabase() {
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func init() {
	checkDatabase()
}
func AddChildWord (word string,
	definition string,
	source string,
	user string,
	parentTextHash []byte) int64 {
	wordId := AddWord(word, definition, source)
AddWordRelation(user, parentTextHash, wordId)

	stmt, err := db.Prepare(``)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec()
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

	return lastId
}
func AddText (contents string, 
	source string) int64 {
	textHash := sha256.Sum256([]byte(contents))

	stmt, err := db.Prepare(`INSERT INTO text (contents, hash, source) 
SELECT 
	? AS contents,
	? AS hash,
	source.id AS source 
FROM source WHERE source.name = ?`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(contents, textHash[:], source)
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

	return lastId
}
func AddWord(word string, 
	definitionText string, 
	source string) int64 {
	definitionTextId := AddText(definitionText, source)

	stmt, err := db.Prepare(`INSERT INTO word (word, definition)
VALUES (?, ?)`)
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

	return lastId
}
func AddSource(source string) int64 {
	

	stmt, err := db.Prepare(`INSERT INTO source (name) VALUES (?)`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(source)
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

	return lastId
}
func AddWordRelation(userEmail string, 
	parentTextHash []byte,
	wordId int64) int64 {
	

	stmt, err := db.Prepare(`INSERT INTO word_relation (user, text, word) 
SELECT 
	user.id AS user,
	text.id AS text,
	? AS word
FROM user 
INNER JOIN text ON text.hash = ?
WHERE user.email = ?`)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(wordId, parentTextHash[:], userEmail)
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

	return lastId
}
