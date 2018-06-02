package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

//Ready : Event triggered when BOT is ready
func Ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	log.Println("Ready!")
	pretty(s.Channel("315197390271414275"))
}
