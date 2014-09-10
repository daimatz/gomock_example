package main

import "github.com/daimatz/gomock_example/twitter_bot"

func main() {
	service := &twitter_bot.TwitterServiceImpl{}
	bot := &twitter_bot.TwitterBotImpl{service}
	bot.EventLoop()
}
