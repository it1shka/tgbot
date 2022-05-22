package setup

import "math/rand"

func RepliesOnAddQuote() []string {
	return []string{
		"Записал мудрость сверху 😌😌😌",
		"Умно сказано! 🥸🥸🥸",
		"Христианская цитата, православная 🥵🥵🥵",
		"Брависсимо ✊✊✊",
		"Хорош 😉😉😉",
		"Теперь это моя любимая цитата 🥰🥰🥰",
	}
}

func RepliesOnGetQuote() []string {
	return []string{
		"Однажды Великий Мудрец 🤓🤓🤓 сказал: ",
		"Мудрость от 🤩🤩🤩 Евгения Хомича 🤩🤩🤩: ",
		"Цитата, которая поднимет 🤗🤗🤗🤗 твой день со дна: ",
		"Лучшее высказывание 😎😎😎",
	}
}

func RandomOnAddQuote() string {
	replies := RepliesOnAddQuote()
	ridx := rand.Intn(len(replies))
	return replies[ridx]
}

func RandomOnGetQuote() string {
	replies := RepliesOnGetQuote()
	ridx := rand.Intn(len(replies))
	return replies[ridx]
}
