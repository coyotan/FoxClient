package FoxClient

import (
	"flag"
	"strings"
	"github.com/bwmarrin/discordgo"
)

type FoxClient struct {
	dg *discordgo.Session
	focusChan []string
}

var (
	Token string
	Mode string
	LogPath string
	
	
	fc FoxClient
	
	version = "0.2.1a"
)



func init() {
	flag.StringVar(&Mode,"m","bot","Used to set the login mode. Options are CLI or bot, default is bot")
	flag.StringVar(&Token,"t","","Used to define the Discord auth token.")
	flag.Parse()

	utils.NewLog(*LogPath)
	utils.Log.Println("PROGRAM INIT STARTED")
	
	if strings.ToLower(Mode) == "bot" {
		utils.Log.Println("Authenticating with a bot token")
		runBot()
	} else if strings.ToLower(Mode) == "cli" {
		utils.Log.Println("Authenticating with a client token")
		runCli()
	} else {
		utils.Log.Panic("ERROR: UNKOWN FLAG, PLEASE SEE DOCUMENTATION")
		fmt.Println("Please check the logs for more information. Ensure that you are using the flags properly")
		os.exit(1)
	}
}

function runCli() {
	
}

function runBot() {
	
}
