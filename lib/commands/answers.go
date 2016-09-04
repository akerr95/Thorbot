package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var tables = []string{
	"(╯°□°）╯︵ ┻━┻",
	"(┛◉Д◉)┛彡┻━┻",
	"(ﾉ≧∇≦)ﾉ ﾐ ┸━┸",
	"(ノಠ益ಠ)ノ彡┻━┻",
	"(╯ರ ~ ರ）╯︵ ┻━┻",
	"(┛ಸ_ಸ)┛彡┻━┻",
	"(ﾉ´･ω･)ﾉ ﾐ ┸━┸",
	"(ノಥ,_｣ಥ)ノ彡┻━┻",
	"(┛✧Д✧))┛彡┻━┻",
}

var flipCount = int(0)

func Answers(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Get the account information.
	u, err := s.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Ignore all messages created by the bot itself
	if m.Author.ID == u.ID {
		return
	}

	// Search for table flips
	if stringInSlice(m.Content, tables) {
		if flipCount == 3 {
			escapeFlip(s, m)
			flipCount = 0
			return
		}

		tableFlip(s, m)
		flipCount++
		return
	}
}

// Respond with the right table flip
func tableFlip(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "┬──┬ ノ( ゜-゜ノ) Don't hate on tables!")
}

func escapeFlip(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "ﾐ┻┻(ﾉ>｡<)ﾉ He's mad!!")
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
