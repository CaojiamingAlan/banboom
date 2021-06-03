package service

import "banboom/mysql"

func TranslateText(text string) string {

	textItem, _ := mysql.SelectEncryptedText(text)
	return textItem.DecryptedText
}
