package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func RunCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	command := args[0]

	switch strings.ToLower(command) {
	case "heartbeat":
		ping(s, m)
	}
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "I'm alive!")
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return
	}
}
