package speech

import (
	"../config"
	"../tools"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

var entryMessages = []string{
	"You people are so petty, and tiny.",
	"You’re big. I’ve fought bigger.",
	"You should know that when you betray me, I will kill you.",
	"Don’t be careful. You could hurt yourself.",
	"Come on, you sons of b*tches! Do you want to live forever?",
	"I bring justice",
	"Injustice anywhere is a threat to justice everywhere.",
	"Peace and justice are two sides of the same coin.",
	"Go ahead. Make my day.",
	"Oh, no... this is Earth... isn't it?",
	"You! What realm is this? Elfheim, Nilfheim?",
	"You dare threaten Thor with such a puny weapon?",
	"This mortal form has grown weak. I need sustenance!",
	"I need a horse!",
}

func Entry(s *discordgo.Session) {

	ticker := time.NewTicker(10 * time.Second)
	savedGuilds := make([]string, 0)

	go func() {
		for {
			select {
			case <-ticker.C:
				for _, element := range s.State.Guilds {
					if !tools.StringInSlice(element.ID, savedGuilds) {
						savedGuilds = append(savedGuilds, element.ID)
						if config.Variables.Application.Debug {
							fmt.Println()
							break
						}
						s.ChannelMessageSend(element.ID, randomizeEntryMessage())
					}
				}
			}
		}
	}()
}

func randomizeEntryMessage() string {
	return entryMessages[rand.Intn(len(entryMessages))]
}
