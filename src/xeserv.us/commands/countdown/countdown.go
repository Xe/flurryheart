package commands

import (
	"fmt"
	"time"

	"xeserv.us/ponyapi"

	"github.com/syfaro/finch"
	"gopkg.in/telegram-bot-api.v4"
)

func init() {
	finch.RegisterCommand(&countdownCommand{})
}

type countdownCommand struct {
	finch.CommandBase
}

func (cmd *countdownCommand) Help() finch.Help {
	return finch.Help{
		Name:        "Countdown",
		Description: "Displays the amount of time unitl the next episode of My Little Pony, Friendship is Magic is aired",
		Example:     "/countdown@@",
		Botfather: [][]string{
			[]string{"countdown", "MLP: FIM episode countdown"},
		},
	}
}

func (cmd *countdownCommand) ShouldExecute(message tgbotapi.Message) bool {
	return finch.SimpleCommand("countdown", message.Text)
}

func (cmd *countdownCommand) Execute(message tgbotapi.Message) error {
	episode, err := ponyapi.Newest()
	if err != nil {
		return err
	}

	now := time.Now()
	then := time.Unix(int64(episode.AirDate), 0)
	diff := then.Sub(now)

	days := int64(diff.Hours() / 24)
	hours := int64(diff.Hours()) % 24
	minutes := int64(diff.Minutes())

	text := fmt.Sprintf(
		"\"%s\" (S%dE%d) will air in %d days, %d hours and %d minutes!",
		episode.Name, episode.Season, episode.Episode,
		days, hours, minutes,
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID

	return cmd.Finch.SendMessage(msg)
}
