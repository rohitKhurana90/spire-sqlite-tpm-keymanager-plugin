package main

import "log"

// import (
//
//	"database/sql"
//	"fmt"
//	"github.com/dixonwille/wmenu"
//	_ "github.com/mattn/go-sqlite3"
//	"log"
//
// )
//
//	func main() {
//		//db := connectToDatabase()
//		db, err := connectToDatabase()
//		checkErr(err)
//
//		menu := wmenu.NewMenu("What would you like to do?")
//
//		menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })
//
//		menu.Option("Add a new Person", 0, true, nil)
//		menu.Option("Find a Person", 1, false, nil)
//		menu.Option("Update a Person's information", 2, false, nil)
//		menu.Option("Delete a person by ID", 3, false, nil)
//		menuerr := menu.Run()
//
//		if menuerr != nil {
//			log.Fatal(menuerr)
//		}
//	}
//
// func handleFunc(db *sql.DB, opts []wmenu.Opt) {
//
//		switch opts[0].Value {
//
//		case 0:
//			//reader := bufio.NewReader(os.Stdin)
//			////fmt.Println("Enter a key to store")
//			////keyIdRead, _ := reader.ReadString('\n')
//			////keyIdRead = strings.TrimSuffix(keyIdRead, "\n")
//
//			//encryptAndSave(db)
//		case 1:
//			//fetchKeysAndDecrypt(db)
//		case 2:
//			fmt.Println("Update a Person's information")
//		case 3:
//			fmt.Println("Deleting a person by ID")
//		case 4:
//			fmt.Println("Quitting application")
//		}
//	}
//
//	/*func fetchKeysAndDecrypt(db *sql.DB) {
//		keys := getKeys(db)
//		for _, key := range keys {
//			fmt.Printf("\n----\nid: %s \nkeyId: %s \nrandomVal: %s \ntypeVal: %s", key.id, key.keyValue, key.randomValue, key.typeValue)
//		}
//		randomValue, _ := hex.DecodeString(keys[0].randomValue)
//		encString := keys[0].keyValue
//		decryptedString, _ := DecryptData(encString, randomValue)
//		fmt.Printf("\n Decrypted string %s", decryptedString)
//	}*/
//
//	/*func encryptAndSave(db *sql.DB) {
//		randomDerivedKey := getRandomDerivedKey()
//		encString, _ := EncryptData("abc", randomDerivedKey)
//		randomDerivedKeyValue := hex.EncodeToString(randomDerivedKey)
//		keyObject := signingKeys{keyValue: encString, randomValue: randomDerivedKeyValue, typeValue: "JWT_SIGNER"}
//		createKeys(db, keyObject)
//	}*/
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
