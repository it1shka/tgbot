package setup

import (
	"fmt"

	"github.com/it1shka/tgbot/database"
	"github.com/lithammer/dedent"
	tb "gopkg.in/tucnak/telebot.v2"
)

func SetupBot(bot *tb.Bot) {

	bot.Handle("/start", func(m *tb.Message) {
		name := m.Sender.FirstName
		resp := fmt.Sprintf(dedent.Dedent(`
			Привет, %s! Если ты фанат Евгения,
			то ты по адресу
			***
			команда /motivate -- получение мотивации
			остальные сообщения -- добавление цитаты в систему
		`), name)
		bot.Send(m.Sender, resp)
	})

	bot.Handle("/motivate", func(m *tb.Message) {
		random_repl := RandomOnGetQuote()
		bot.Send(m.Sender, random_repl)
		random_quote := database.GetRandomQuoteString()
		bot.Send(m.Sender, random_quote)
	})

	bot.Handle("/all", func(m *tb.Message) {
		allstr := database.GetAllQuotesString()
		bot.Send(m.Sender, allstr)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		quote := m.Text
		author := m.Sender.Username
		database.AddQuote(quote, author)
		random_repl := RandomOnAddQuote()
		bot.Send(m.Sender, random_repl)
	})

}
