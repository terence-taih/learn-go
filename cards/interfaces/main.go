package main

import "fmt"

type bot interface {
	getGetting() string
}

type workspace struct {
	name string
	bot
	// wsBot bot
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	englishWS := workspace{
		bot:  englishBot{},
		name: "English Workspace",
	}
	englishWS.welcome()

	spanishhWS := workspace{
		bot:  englishBot{},
		name: "English Workspace",
	}
	spanishhWS.welcome()

}

func (ws workspace) welcome() {
	fmt.Println(ws.name, ws.bot.getGetting())
}

func (eb englishBot) getGetting() string {
	return "Hello"
}

func (sb spanishBot) getGetting() string {
	return "Hallo"
}
