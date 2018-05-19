package main

import (
	"github.com/TheBoxFox/FoxClient/utils"
	"strings"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"os"
)

type FoxClient struct {
	dg *discordgo.Session

	Core utils.Conf
}


var (
	LogPath = "./connection.log"

	fc FoxClient
	
	version = "0.2.1a"
)



func init() {
	fc.Core = utils.Parse("./config.json")

	utils.NewLog(LogPath)
	utils.Log.Println("PROGRAM INIT STARTED ",version)

	utils.Log.Println(fc.Core.Client.Mode)
	if strings.ToLower(fc.Core.Client.Mode) == "cli" {
		dg, err := discordgo.New(fc.Core.Client.Token)
		fc.dg = dg

		if err != nil {
			utils.Log.Println("Error Authenticating with Discord as client")
			ErrFunc(err)
			os.Exit(1)
		}

		runCli()

		for i:=0; i > len(fc.Core.Watchman.Notif); i++ {
			fmt.Println(fc.Core.Watchman.Notif[i].Name)
		}
	} else if strings.ToLower(fc.Core.Client.Mode) == "bot" {

		dg, err := discordgo.New("bot "+fc.Core.Client.Token)
		fc.dg = dg

		if err != nil {
			utils.Log.Println("Error Authenticating with Discord as bot")
			ErrFunc(err)
			os.Exit(1)
		}

		runBot()
	} else {
		utils.Log.Panic("ERROR: UNKNOWN VALUE, PLEASE SEE DOCUMENTATION")
		utils.Log.Println("Invalid authentication mode: "+fc.Core.Client.Mode)
		fmt.Println("Please check the logs for more information. Ensure that you are using the flags properly")
		os.Exit(1)
	}
}
func main () {}

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