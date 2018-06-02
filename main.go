package main

import (
	"fmt"
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

	shell.AddCmd(&ishell.Cmd{
		Name: "connect",
		Help: "connect to a bot using token [usage: login {token}]",
		Func: func(c *ishell.Context) {
			token := strings.Join(c.Args, " ")
			discord, err := discordgo.New("Bot " + token)
			if err != nil {
				log.Fatal(err)
				return
			}

			discord.AddHandler(Ready)

			err = discord.Open()
			if err != nil {
				fmt.Println("Error opening Discord session: ", err)
			}

		},
	})

	shell.Run()
}
