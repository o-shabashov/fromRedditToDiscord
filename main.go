package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/jzelinskie/geddit"
)

func main() {
	err := godotenv.Load()
	CheckErr(err)

	o, err := geddit.NewOAuthSession(
		os.Getenv("REDDIT_CLIENT_ID"),
		os.Getenv("REDDIT_CLIENT_SECRET"),
		"Testing OAuth Bot by u/my_user v0.1 see source https://github.com/jzelinskie/geddit",
		os.Getenv("REDDIT_REDIRECT_URL"),
	)
	CheckErr(err)

	// Create new auth token for confidential clients (personal scripts/apps).
	err = o.LoginAuth(os.Getenv("REDDIT_USERNAME"), os.Getenv("REDDIT_PASSWORD"))
	CheckErr(err)

	// Ready to make API calls!
	subOpts := geddit.ListingOptions{
		Limit: 25,
	}
	// Get our own personal frontpage
	submissions, _ := o.Frontpage(geddit.DefaultPopularity, subOpts)

	// Connect to Discord
	dg, err := discordgo.New(os.Getenv("DISCORD_BOT_ID"))
	CheckErr(err)

	// Print title and author of each submission
	for _, s := range submissions {
		_, err := dg.ChannelMessageSend(os.Getenv("DISCORD_CHANNEL_ID"), fmt.Sprintf("%s\n%s\n", s.Title, s.URL))
		CheckErr(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
