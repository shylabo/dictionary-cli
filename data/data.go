package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	CreateTableSQL := `CREATE TABLE IF NOT EXISTS test_table (
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
		);`

	statement, err := db.Prepare(CreateTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("test_table created!")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO test_table (word, definition, category)
	VALUES(?, ?, ?)`

	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted new note success")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM test_table ORDER BY word")
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string

		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "]", word, "-", definition)
	}
}
