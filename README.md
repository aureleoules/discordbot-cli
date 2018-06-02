<p align="center">
  <img src="https://discordapp.com/assets/f8389ca1a741a115313bede9ac02e2c0.svg" alt="" width=150 height=150>

  <h1 align="center">Discord BOT - CLI</h1>

  <p align="center">
    Control a <strong>Discord BOT</strong> using this simple <strong>Command Line Tool</strong> 
  </p>
</p>

## Usage
In order to make this tool work, you have two options: 
* Download the source of this repository then compile it with :
```
go run *go
```
or on Windows
```
go build && discordbot-cli.exe
```
* Download the latest release [here](https://github.com/aureleoules/discordbot-cli/releases)

Multiple commands are available :
* **clear** : clear the screen
*  **config** : Configure your config.json 
_usage_: config -github {key}
*  **connect** : connect to a bot using token 
_usage_: connect {token}
*  **exit** : exit the program
*  **help** : display help
*  **invite** : creates an invite to the server
*  **message** : send message to channel or user 
_usage_: message -u {username} -d {userDiscriminator} -cid channelId -c channelName "message"
*  **roles** : shows roles of user 
_usage_: roles -u {username} -d {userDiscriminator}
*  **tokens** : look for bot tokens on github 
_usage_: tokens {page}
*  **user** : Display informations of an user 
_usage_: user -u {username} -d {discriminator}
*  **users** : shows all users

## Author
[Aurèle Oulès](http://aurele.oules.com)