package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"       // for accessing version 1.1 of the Twitter API
	"github.com/dghubble/go-twitter/twitter" //  for providing a Client for the Twitter API
	"github.com/dghubble/oauth1"             // authorization flow and provides an http.Client
	"github.com/sirupsen/logrus"             // API compatible structured logger
)

// Credentials stores all of our access/consumer tokens
// and secret keys needed for authentication for twitter REST API
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
	// authentication of Twitter API for anaconda
	anaconda.SetConsumerKey(creds.ConsumerKey)
	anaconda.SetConsumerSecret(creds.ConsumerSecret)
	api := anaconda.NewTwitterApi(creds.AccessToken, creds.AccessTokenSecret)

	// setup log for items recieved by anaconda
	log := &logger{logrus.New()}
	api.SetLogger(log)

	// stream has access all public tweets from Twitter
	// then filter the tweets based on a specific phrase or hashtag
	stream := api.PublicStreamFilter(url.Values{
		// here we're filtering for #coding
		"track": []string{"lol"},
	})

	// stop at the end of program
	defer stream.Stop()

	// access channel(C)
	for v := range stream.C {
		// check that the value recieved from channel C is of type anaconda.Tweet
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("received unexpected value of type %T", v)
			continue
		}

		// check if a tweet has already been retweeted
		// if it has then do not retweet and continue stream
		if t.RetweetedStatus != nil {
			continue
		}

		// retweet specific tweet(Id)
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
	// get credentials from environment
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	// fmt.Printf("%+v\n", creds) // verify personal twitter information in terminal

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// stream phrase or hashtag of choice
	setupTwitterStream(&creds)

	// search tweets to retweet determined by searchParams
	searchParams := &twitter.SearchTweetParams{
		Query:      "#coding",
		Count:      3,
		ResultType: "trending",
		Lang:       "en",
	}

	// pass searchParams to search function
	searchResult, _, _ := client.Search.Tweets(searchParams)

	// retweet
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
}

type logger struct {
	*logrus.Logger
}

// critical methods for logger -- do not remove
func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }

// shoutouts: https://twitter.com/Elliot_F
