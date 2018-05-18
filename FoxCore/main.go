package FoxCore

import (
	"../utils"

	"flag"
	"strings"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
)

type FoxClient struct {
	dg *discordgo.Session

}

type WatchDog struct {
	channel *discordgo.Channel //Watch for messages in a specific channel
	guild *discordgo.Guild //Watch all content from a specific server
	user *discordgo.User //Watch for all messages from a specific user
	role *discordgo.GuildRole //Watch for messages from a specific role
	keyword string //watch for any messages containing a specific keyword

	color string //Color making for watched events
	priority string //Priority control for watched events
}

var (
	Token string
	Mode string
	LogPath string
	notifiers []WatchDog

	fc FoxClient
	
	version = "0.2.1a"
)



func init() {
	flag.StringVar(&Mode,"m","bot","Used to set the login mode. Options are CLI or bot, default is bot")
	flag.StringVar(&Token,"t","","Used to define the Discord auth token.")
	flag.Parse()

	utils.NewLog(LogPath)
	utils.Log.Println("PROGRAM INIT STARTED %s",version)
	
	if strings.ToLower(Mode) == "bot" {
		utils.Log.Println("Authenticating with a bot token")
		Token = "bot " + Token

		dg, err := discordgo.New(Token)//Create Discord Session as bot
		fc.dg = dg

		if err != nil {
			ErrFunc(err)
		}

		runBot()

	} else if strings.ToLower(Mode) == "cli" {
		utils.Log.Println("Authenticating with a client token")

		dg, err := discordgo.New(Token) //Create Discord Session as client
		fc.dg = dg

		if err != nil {
			ErrFunc(err)
		}

		runCli()

	} else {
		utils.Log.Panic("ERROR: UNKOWN FLAG, PLEASE SEE DOCUMENTATION")
		fmt.Println("Please check the logs for more information. Ensure that you are using the flags properly")
		os.Exit(1)
	}
}

func runCli() {
	utils.Log.Println("RunCli Start")

	fc.dg.AddHandler(Ready)
}



func runBot() {
	utils.Log.Println("RunBot Start")
}

func ErrFunc(err error) {
	utils.Log.Println("Unexpected error occured, please see error dump below.")
	utils.Log.Panic(err)
}

//Create basic handlers to start the bot

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.Log.Println("Discord Ready Message Recieved. Username: %s User ID: %s",r.User.Username, r.User.ID)
}