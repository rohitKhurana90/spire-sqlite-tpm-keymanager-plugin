package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func connectToDatabase() (*sql.DB, error) {
	//Connect to the database
	err := createEmptyFile("./signing_keys.db")
	if err != nil {
		log.Println("Error creating file")
		log.Println(err)
		return nil, err
	}

	// Connect to the database
	db, err := sql.Open("sqlite3", "./signing_keys.db")
	checkErr(err)
	log.Println("DB Created")

	// Change file permissions
	err = setFilePermissions("./signing_keys.db", 0777) // Adjust the permission bits as needed
	if err != nil {
		log.Println("Error setting file permissions")
		log.Println(err)
		return nil, err
	} // Create the table if it doesn't exist
	createTableQuery := "CREATE TABLE IF NOT EXISTS signing_keys (id INTEGER PRIMARY KEY AUTOINCREMENT,key_val TEXT, type_val TEXT, random_val TEXT);"
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Println("Error during table creation")
		log.Println(err)
	}
	checkErr(err)
	fmt.Println("Connected to database successfully.")
	return db, nil
}

func createEmptyFile(filePath string) error {
	// Check if the file already exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// File does not exist, create it
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func setFilePermissions(filePath string, permissions os.FileMode) error {
	err := os.Chmod(filePath, permissions)
	if err != nil {
		return err
	}
	return nil
}

func createKeys(db *sql.DB, keys signingKeys) {

	stmt, _ := db.Prepare("INSERT INTO signing_keys (key_val, type_val, random_val) VALUES (?,?,?)")
	stmt.Exec(keys.keyValue, keys.typeValue, keys.randomValue)
	defer stmt.Close()

	fmt.Printf("Added %v \n", keys.keyValue)
}

func getKeys(db *sql.DB) []signingKeys {

	rows, err := db.Query("SELECT * FROM signing_keys")

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	response := make([]signingKeys, 0)

	for rows.Next() {
		key := signingKeys{}
		err = rows.Scan(&key.id, &key.keyValue, &key.typeValue, &key.randomValue)
		if err != nil {
			log.Fatal(err)
		}

		response = append(response, key)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return response
}
