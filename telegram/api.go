package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	lua "github.com/yuin/gopher-lua"
)

const (
	// max idle connections
	MaxIdleConns = 1
	// max open connections
	MaxOpenConns = 1
)

type lTgBot interface {
	constructor(string) (lTgBot, error)
	getBot() *tgbotapi.BotAPI
}

var tgBot = new(luaTgBot)

func New(L *lua.LState) int {

	token := L.CheckString(1)
	result, err := tgBot.constructor(token)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	ud := L.NewUserData()
	ud.Value = result
	L.SetMetatable(ud, L.GetTypeMetatable(`telegram_ud`))
	L.Push(ud)
	return 1
}
func checkBot(L *lua.LState, n int) lTgBot {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(lTgBot); ok {
		return v
	}
	L.ArgError(n, "telegram expected")
	return nil
}

func Send(L *lua.LState) int {
	botInterface := checkBot(L, 1)
	chatId := L.CheckInt64(2)
	msgText := L.CheckString(3)
	msg := tgbotapi.NewMessage(chatId, msgText)

	bot := botInterface.getBot()
	_, err := bot.Send(msg)
	if err != nil {
		L.Push(lua.LString(err.Error()))
		return 1
	}
	return 1
}
