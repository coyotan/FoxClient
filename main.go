package FoxClient

import (
	"flag"
	"github.com/bwmarrin/discordgo"
)

type FoxClient struct {
	dg *discordgo.Session
	focusChan []string
}

var (
	Token string
	Mode string

	fc FoxClient
)



func init() {
	flag.StringVar(&Mode,"m","bot","Used to set the login mode. Options are CLI or bot, default is bot")
	flag.StringVar(&Token,"t","","Used to define the Discord auth token.")
	flag.Parse()
//Just a test to see if it syncs.


}
