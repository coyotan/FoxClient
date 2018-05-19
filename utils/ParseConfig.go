package utils

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type WatchDogs struct {
	Name string `json:"Name"`
	Channel *discordgo.Channel `json:"Channel"`//Watch for messages in a specific channel
	Guild *discordgo.Guild  `json:"Guid"`//Watch all content from a specific server
	User *discordgo.User `json:"User"`//Watch for all messages from a specific user
	Role *discordgo.GuildRole `json:"Role"`//Watch for messages from a specific role
	Keyword string `json:"Keyword"`//watch for any messages containing a specific keyword

	Color string `json:"Color"`//Color making for watched events
	Priority string `json:"Priority"`//Priority control for watched events
}

//type WatchDog struct {
//	Notif []WatchDogs `json:"WatchDogs"`
//}

type core struct {
	Token string `json:"Token"`
	Mode string `json:"Mode"`
}

type Conf struct {
	Client core
	WatchDog []WatchDogs
}


func Parse(filename string) Conf {

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("There was a critical error opening the json.")
		os.Exit(1)
	}

	defer file.Close()

	byteVal, _ := ioutil.ReadAll(file)

	var client Conf

	json.Unmarshal(byteVal,&client)
	Log.Println(client)
	return client
}