package service

import (
	"banboom/mysql"
	"log"
	"regexp"
)
func TranslateText(text string) string {
	r := regexp.MustCompile(`[a-zA-Z]+`)
	v := r.FindAllStringIndex(text, -1)
	result := ""
	prev := 0
	for i, s := range v {
		result += text[prev:s[0]]
		prev = s[1]
		log.Println(i, s)
		textItem, _ := mysql.SelectEncryptedText(text[s[0]:s[1]])
		result += textItem.DecryptedText
	}
	return result
}
