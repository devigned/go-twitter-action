package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	flag "github.com/spf13/pflag"
)

var (
	message, apiKey, apiKeySecret, accessToken, accessTokenSecret string
)

func main() {
	parseAndValidateInput()
	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	if os.Getenv("DRY_RUN") != "true" {
		_, _, err := client.Statuses.Update(message, nil)
		if err != nil {
			PrintOutput("errorMessage", err.Error())
			log.Fatalf("status update error: %q", err)
		}
	}
}

func parseAndValidateInput() {
	flag.StringVar(&message, "message", "", "message you'd like to send to twitter")
	flag.StringVar(&apiKey, "apiKey", "", "twitter api key")
	flag.StringVar(&apiKeySecret, "apiKeySecret", "", "twitter api key secret")
	flag.StringVar(&accessToken, "accessToken", "", "twitter access token")
	flag.StringVar(&accessTokenSecret, "accessTokenSecret", "", "twitter access token secret")
	flag.Parse()

	if message == "" {
		log.Fatal("--message can't be empty")
	}

	if apiKey == "" {
		log.Fatal("--apiKey can't be empty")
	}

	if apiKeySecret == "" {
		log.Fatal("--apiKeySecret can't be empty")
	}

	if accessToken == "" {
		log.Fatal("--accessToken can't be empty")
	}

	if accessTokenSecret == "" {
		log.Fatal("--accessTokenSecret can't be empty")
	}
}

func PrintOutput(key, message string) {
	fmt.Printf("::set-output name=%s::%s\n", key, message)
}
