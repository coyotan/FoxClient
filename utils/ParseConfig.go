package utils

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type WatchDogCell struct {
	Name string `json:"Name"`
	Channel *discordgo.Channel `json:"Channel"`//Watch for messages in a specific channel
	guild *discordgo.Guild  `json:"Guid"`//Watch all content from a specific server
	user *discordgo.User `json:"User"`//Watch for all messages from a specific user
	role *discordgo.GuildRole `json:"Role"`//Watch for messages from a specific role
	keyword string `json:"Keyword"`//watch for any messages containing a specific keyword

	color string `json:"Color"`//Color making for watched events
	priority string `json:"Priority"`//Priority control for watched events
}

type WatchDog struct {
	Notif []WatchDogCell `json:"WatchDogs"`
}

type core struct {
	Token string `json:"Token"`
	Mode string `json:"Mode"`
}

type Conf struct {
	Client core
	Watchman WatchDog
}

func parse(name string) Conf {

	file, err := os.Open(name)

	if err != nil {
		fmt.Println("There was a critical error opening the json.")
		os.Exit(1)
	}

	defer file.Close()

	byteVal, _ := ioutil.ReadAll(file)

	var client Conf

	json.Unmarshal(byteVal,&client)

	return client
}