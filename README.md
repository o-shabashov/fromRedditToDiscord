# From Reddit to Discord

[![BCH compliance](https://bettercodehub.com/edge/badge/o-shabashov/fromRedditToDiscord?branch=master)](https://bettercodehub.com/)

Grab images (25 by default) from your Reddit account frontpage and post it to Discord channel.

Using OAuth Reddit and Discord Bot.

## Install dependencies
```shell
glide install
```

## Edit .env file

* `REDDIT_CLIENT_ID` and `REDDIT_CLIENT_SECRET` - https://github.com/reddit/reddit/wiki/OAuth2

* `REDDIT_USERNAME` and `REDDIT_PASSWORD` - your Reddit account credentials.

* `REDDIT_REDIRECT_URL` - your IP, accessibly from internet.

* `DISCORD_BOT_ID` - https://github.com/reactiflux/discord-irc/wiki/Creating-a-discord-bot-&-getting-a-token

* `DISCORD_CHANNEL_ID` - https://github.com/Chikachi/DiscordIntegration/wiki/How-to-get-a-token-and-channel-ID-for-Discord#get-the-channel-id-of-the-discord-text-channel

## Run

```shell
go run main.go
```
