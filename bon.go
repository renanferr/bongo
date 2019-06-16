package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Session : discord bot session
var Session, _ = discordgo.New()

func main() {
	log.Println("Starting up...")
	var err error
	Session.Token = "Bot " + os.Getenv("DISCORD_BOT_TOKEN")
	if Session.Token == "Bot " {
		log.Println("You must provide a Discord authentication token.")
		os.Exit(1)
	}

	err = Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	Session.AddHandler(messageCreate)

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "join" {
		var err error
		var guild *discordgo.Guild

		guild, err = s.Guild(m.GuildID)
		if err != nil {
			log.Printf("%v", err)
			return
		}
    authorVoiceState := filter(guild.VoiceStates, m)
    if authorVoiceState == nil {
      s.ChannelMessageSend(m.ChannelID, "Entre num canal de voz antes de solicitar o join!")
      return
    }

		s.ChannelVoiceJoin(guild.ID, authorVoiceState.ChannelID, false, false)
	}
	
}

func filter(coll []*discordgo.VoiceState, m *discordgo.MessageCreate) *discordgo.VoiceState {
    var authorVoiceState *discordgo.VoiceState
    for _, vs := range coll {
        if vs.UserID == m.Author.ID {
            authorVoiceState = vs
        }
    }
    return authorVoiceState
}