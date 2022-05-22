package database

import (
	"fmt"
	"strings"

	"github.com/lithammer/dedent"
)

type Quote struct {
	Id      uint   `gorm:"primaryKey;autoIncrement;notNull" json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func AddQuote(content, author string) {
	quote := Quote{
		Content: content,
		Author:  author,
	}
	DB.Create(&quote)
}

func GetRandomQuote() Quote {
	var result Quote
	db := DB.Raw("select * from quotes order by random() limit 1;")
	db.Scan(&result)
	return result
}

func FormatQuote(quote Quote) string {
	return fmt.Sprintf("%s\n\nДобавил @%s", quote.Content, quote.Author)
}

func GetRandomQuoteString() string {
	quote := GetRandomQuote()
	return FormatQuote(quote)
}

func GetAllQuotesString() string {
	var quotes []Quote
	DB.Find(&quotes)
	var result strings.Builder
	for idx, each := range quotes {
		strquote := dedent.Dedent(fmt.Sprintf(`
			%d. От @%s:

			%s
		`, idx+1, each.Author, each.Content))
		result.WriteString(strquote)
	}
	return result.String()
}
