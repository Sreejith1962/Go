package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7566477146833-7553907808914-qsiUxSzJnYhWpbtRRhwHV8yd")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A07GNEP33H7-7553836641810-9b156eaa832ee79c91e8fb049459696bf244c256209a55fce2a8373836ee6951")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	eg := []string{"my yob is 2020"}

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Examples:    eg,
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			year := r.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2024 - yob
			re := fmt.Sprintf("Age is : %d", age)
			w.Reply(re)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events ")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Event)
		fmt.Println(event.Parameters)
		fmt.Println(event.Command)
		fmt.Println()
	}
}
