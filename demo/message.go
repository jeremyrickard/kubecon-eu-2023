package main

type Message struct {
	Sender string `mapstructure:"sender"`
	Text   string `mapstructure:"text"`
}
