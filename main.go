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
			utils.Log.Println("Error authenticating with Discord as client")
			ErrFunc(err)
			os.Exit(2)
		}

		err = fc.dg.Open()
		
		if err != nil {
			utils.Log.Println("Error opening new discord session.")
			ErrFunc(err)
			os.Exit(3)
		}
		
		runCli()


	} else if strings.ToLower(fc.Core.Client.Mode) == "bot" {

		dg, err := discordgo.New("Bot "+fc.Core.Client.Token)
		fc.dg = dg

		if err != nil {
			utils.Log.Println("Error authenticating with Discord as bot")
			ErrFunc(err)
			os.Exit(1)
		}
		
		err = fc.dg.Open()
		
		if err != nil {
			utils.Log.Println("Error opening new Discord bot session. Are you sure you put in the correct token?")
			ErrFunc(err)
			os.Exit(3)
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
	utils.Log.Println("Launching Client in CLI mode.")

	fc.dg.AddHandler(Ready)
}



func runBot() {
	utils.Log.Println("Launching Client in Bot mode.")
}

func ErrFunc(err error) {
	utils.Log.Println("BEGIN ERROR DUMP")
	utils.Log.Panic(err)
}

//Create basic handlers to start the bot

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	utils.Log.Println("Discord Ready Message Recieved. Username:",r.User.Username," User ID:", r.User.ID)
	fc.local = r
}

/*
Error numbers are as follows
1: Configuration error
2: Discord authentication error
3: Discord session based error
*/
