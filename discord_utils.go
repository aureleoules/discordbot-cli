package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//FindFirstGuild : Returns first Guild
func FindFirstGuild(s *discordgo.Session) *discordgo.Guild {
	return s.State.Guilds[0]
}

//FindFirstChannel : Look for first channel
func FindFirstChannel(s *discordgo.Session) *discordgo.Channel {
	guildChannels := FindFirstGuild(s).Channels
	channelID := guildChannels[0].ID
	channel, _ := s.Channel(channelID)
	return channel
}

//FindUserByUsername : find user by username
func FindUserByUsername(s *discordgo.Session, username string) *discordgo.Member {
	members := FindFirstGuild(s).Members
	for i := range members {
		member := members[i]
		if strings.ToLower(member.User.Username) == strings.ToLower(username) {
			return member
		}
	}
	return nil
}

//FindUserByDiscriminator : find user by discriminator
func FindUserByDiscriminator(s *discordgo.Session, discriminator string) *discordgo.Member {
	members := FindFirstGuild(s).Members
	for i := range members {
		member := members[i]
		if member.User.Discriminator == discriminator {
			return member
		}
	}
	return nil
}

//FindAllUsers : Find all users
func FindAllUsers(s *discordgo.Session) []*discordgo.Member {
	guild := FindFirstGuild(s)
	return guild.Members
}

//FindChannelByID : find channel by id
func FindChannelByID(s *discordgo.Session, id string) *discordgo.Channel {
	channels := FindFirstGuild(s).Channels
	for i := range channels {
		channel := channels[i]
		if channel.ID == id {
			return channel
		}
	}
	return nil
}

//FindChannelByName : find channel by name
func FindChannelByName(s *discordgo.Session, name string) *discordgo.Channel {
	channels := FindFirstGuild(s).Channels
	for i := range channels {
		channel := channels[i]
		if strings.ToLower(channel.Name) == strings.ToLower(name) {
			return channel
		}
	}
	return nil
}

//CreateInvite : Creates an invite to a specific channel
func CreateInvite(s *discordgo.Session, channel *discordgo.Channel) (string, error) {
	invite, err := s.ChannelInviteCreate(channel.ID, discordgo.Invite{MaxUses: 0, MaxAge: 0, Temporary: true})
	if err != nil {
		log.Println(err)
		return "", err
	}
	url := "http://discord.gg/" + invite.Code
	return url, nil
}

//CreateChannel : Creates a discord channel
func CreateChannel(s *discordgo.Session, id string, name string, channelType string) (*discordgo.Channel, error) {
	channel, err := s.GuildChannelCreate(id, name, channelType)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return channel, nil
}
