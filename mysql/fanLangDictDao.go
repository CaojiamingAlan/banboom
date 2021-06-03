package mysql

import (
	"errors"
	"log"
)

type TextItem struct {
	ID            int    `json:"id"`
	EncryptedText string `json:"encrypted_text"`
	DecryptedText string `json:"decrypted_text"`
}

func SelectEncryptedText(text string) (TextItem, error) {
	log.Printf("Enter selectEncryptedText")
	results, err := Db.Query("SELECT * FROM fan_lang_dict WHERE encrypted_text = '" + text + "'")
	if err != nil {
		panic(err.Error())
	}
	var textItem TextItem
	textItem.DecryptedText = text
	for results.Next() {
		err = results.Scan(&textItem.ID, &textItem.EncryptedText, &textItem.DecryptedText)
		if err != nil {
			log.Fatal(err.Error()) // proper error handling instead of panic in your app
			break
		}
		log.Printf("%d %s %s", textItem.ID, textItem.EncryptedText, textItem.DecryptedText)
		err = results.Close()
		if err != nil {
			log.Fatal(err.Error()) // proper error handling instead of panic in your app
			break
		}
		return textItem, nil
	}
	log.Print("text '" + text + "' not in dict")
	return textItem, errors.New("text '" + text + "' not in dict")
}