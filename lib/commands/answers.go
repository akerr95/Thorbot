package commands

import (
	"../tools"
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
	if tools.StringInSlice(m.Content, tables) {
		tableFlip(s, m)
		return
	}

	message := checkQuote(m.Content)

	if message != "" {
		s.ChannelMessageSend(m.ChannelID, message)
	}
}

// Respond with the right table flip
func tableFlip(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "┻━┻ ︵ヽ(`Д´)ﾉ︵ ┻━┻ You! This is an act of war, I shall protect this realm!")
}

func checkQuote(a string) string {

	quotes := []struct {
		quote    string
		response string
	}{
		{"You are a vain, greedy, cruel boy!", "And you are an old man and a fool!"},
		{"Describe exactly what happened to you last night.", "Your ancestors called it magic... *skims through a book on Norse mythology* ...but you call it science. I come from a land where they are one and the same."},
		{"Keeping the bridge open would unleash the full power of the Bifrost and destroy Jotunheim, with you on it.", "I have no plans to die today."},
		{"We don't have horses. Just dogs, cats, birds.", "Then give me one of those large enough to ride."},
		{"NO! I'll die a warrior's death! Stories will be told of this day!", "Live, and tell those stories yourself!"},
		{"What happened?", "We're fine! We drank, we fought - we made our ancestors proud!"},
	}

	for _, command := range quotes {
		if command.quote == a {
			return command.response
		}
	}

	return ""
}
