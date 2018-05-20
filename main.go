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

	local *discordgo.Ready

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

	utils.Log.Println(fc.Core.Client.Mode,"mode selected by json configuration")
	//If there is no token supplied, error so that the system won't attempt to authenticate.
	if fc.Core.Client.Token == "" {
		utils.Log.Panic("No token supplied, Discord authentication will fail")
		os.Exit(1)
	}

	if strings.ToLower(fc.Core.Client.Mode) == "cli" {
		dg, err := discordgo.New(fc.Core.Client.Token)
		fc.dg = dg

		if err != nil {
			utils.Log.Println("Error Authenticating with Discord as client")
			ErrFunc(err)
			os.Exit(1)
		}

		runCli()


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
	utils.Log.Println("Discord Ready Message Recieved. Username:",r.User.Username," User ID:", r.User.ID)
	fc.local = r
}