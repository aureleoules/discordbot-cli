package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

//Ready : Event triggered when BOT is ready
func Ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Println(s.State.User.Username + "is ready!")
}
