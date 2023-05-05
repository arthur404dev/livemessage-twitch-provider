package main

import (
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	TwitchUsername   string `yaml:"twitch_username"`
	TwitchOAuthToken string `yaml:"twitch_oauth_token"`
	ChannelName      string `yaml:"channel_name"`
}

func main() {
	var config Config
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config.yaml: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config.yaml: %v", err)
	}

	twitchClient := twitch.NewClient(config.TwitchUsername, config.TwitchOAuthToken)

	twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		log.Printf("%s: %s\n", message.User.DisplayName, message.Message)
	})

	twitchClient.OnConnect(func() {
		log.Printf("Connected to Twitch chat")
		twitchClient.Join(config.ChannelName)
	})

	err = twitchClient.Connect()
	if err != nil {
		log.Fatalf("Error connecting to Twitch chat: %v", err)
	}
}
