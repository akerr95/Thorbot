package main


import (
	"./lib/commands"
	"./lib/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)


func main() {

	// Create a new Discord session using the provided login information.
	discord, err := discordgo.New(
		config.Variables.Application.Email,
		config.Variables.Application.Password,
		config.Variables.Application.Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register answers
	discord.AddHandler(commands.Answers)

	// Open the websocket and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	<-make(chan struct{})
	return
}