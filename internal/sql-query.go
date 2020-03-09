package internal

import (
	"log"
)

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

