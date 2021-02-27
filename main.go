package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sirupsen/logrus"
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// getClient is a helper function that will return a twitter client
// that we can subsequently use to send tweets, or to stream new tweets
// this will take in a pointer to a Credential struct which will contain
// everything needed to authenticate and return a pointer to a twitter Client
// or an error
func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials to Twitter
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// retrieve user and verify if the credentials are successful
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	// log credentials
	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func setupTwitterStream(creds *Credentials) {
	anaconda.SetConsumerKey(creds.ConsumerKey)
	anaconda.SetConsumerSecret(creds.ConsumerSecret)
	api := anaconda.NewTwitterApi(creds.AccessToken, creds.AccessTokenSecret)
	log := &logger{logrus.New()}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"love"},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("received unexpected value of type %T", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}

		_, err := api.Retweet(t.Id, false)
		if err != nil {
			log.Errorf("could not retweet %d: %v", t.Id, err)
			continue
		}
		log.Infof("retweeted %d", t.Id)
	}
}

func main() {
	// retrieve a new Twitter client
	fmt.Println("Go-Twitter Bot v0.01")
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// fmt.Printf("%+v\n", client)

	// stream phrase or hashtag of choice
	// setupTwitterStream(&creds)

	// Search tweets to retweet
	searchParams := &twitter.SearchTweetParams{
		Query:      "software",
		Count:      5,
		ResultType: "trending",
		Lang:       "en",
	}

	searchResult, _, _ := client.Search.Tweets(searchParams)

	// Retweet
	for _, tweet := range searchResult.Statuses {
		tweetID := tweet.ID
		client.Statuses.Retweet(tweetID, &twitter.StatusRetweetParams{})
		if err != nil {
			log.Println(err)
		}
		// log.Printf("%+v\n", resp)
		log.Printf("%+v\n", tweet)

		fmt.Printf("RETWEETED: %+v\n", tweet.Text)
	}

	// testing status update
	// tweet, resp, err := client.Statuses.Update("Hello world!!", nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("%+v\n", resp)
	// log.Printf("%+v\n", tweet)

}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }

// shoutouts: https://twitter.com/Elliot_F
