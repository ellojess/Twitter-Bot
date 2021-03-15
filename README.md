# Twitter-Bot ![Go Report Card](https://goreportcard.com/badge/github.com/ellojess/Twitter-Bot)

<!-- PROJECT LOGO -->
  <p align="center">
    <br />
    <br />
    <a href="https://drive.google.com/file/d/1XTezAZvB1EM4Gb54J7O92yTlc_F8Z0rw/view?usp=sharing">View Demo & Code Walkthrough</a>
    ·
    <a href="https://github.com/ellojess/Twitter-Bot/issues">Report Bug</a>
    ·
    <a href="https://github.com/ellojess/Twitter-Bot/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!--  [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

This Twitter Bot automates the tracking and retweeting process and was built for users who want their feed to keep active with the latest tweets of their defined hashtag or phrase.

Simply choose between an open stream (for unlimited and continuous retweets) or a set number of retweets via defined search parameters. 

### Built With
* [Go / GoLang](https://getbootstrap.com)
* [Twitter API](https://jquery.com)
* [Twitter App](https://laravel.com)



<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple example steps.

### Prerequisites

- Twitter Developers Account 
- Twitter Account


### Installation

1. Clone the repo
   ```sh
   git clone git@github.com:ellojess/Twitter-Bot.git
   ```
2. Get your free credentials from Twitter at [https://developer.twitter.com/en/dashboard](https://example.com)

    Use the credentials from the Developers Account to get the following values. 

    Replace VALUE with your personal credentials and declare them locally in your terminal 

    ```sh
    $ export CONSUMER_KEY=VALUE

    $ export CONSUMER_SECRET=VALUE

    $ export ACCESS_TOKEN=VALUE

    $ export ACCESS_TOKEN_SECRET=VALUE
    ```
3. To run the bot
   ```sh
   $ go build && go run ./main.go
   ```

<!-- USAGE EXAMPLES -->
## Usage

#### Case for Stream 
Lets say it's the FIFA World Cup and you want your Twitter to be updated with the latest tweets at all times for the event. Without having to take your eyes off the screen, Twitter Bot will manage retweets of the event for you. All you need to do is update the 


```go 
  stream := api.PublicStreamFilter(url.Values{
    "track": []string{"#worldcup"},
  })
```

This code snippet defines a stream to track `#worldcup`

#### Case for Set Number of Retweets:
Maybe you haven't retweeted in awhile and you want to quickly post something to keep your Twitter account active. Twitter Bot has the option to retweet a specified number of tweets in a second. 

```go
  searchParams := &twitter.SearchTweetParams{
    Query:      "love",
    Count:      3,
    ResultType: "trending",
    Lang:       "en",
  }
```

This code snippet defines the search parameters for the bot to find `love` in Twitter's platform and returns 3 english posts from `trending`

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/ellojess/Twitter-Bot/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request using this [template](https://github.com/embeddedartistry/templates/blob/master/oss_docs/PULL_REQUEST_TEMPLATE.md)



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Jessica Trinh - [@ellojesss](https://twitter.com/ellojesss) - jtjessicatrinh@gmail.com

Project Link: [https://github.com/ellojess/Twitter-Bot](https://github.com/ellojess/Twitter-Bot)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Anaconda](https://pkg.go.dev/github.com/ChimeraCoder/anaconda)
* [Go-Twitter](https://pkg.go.dev/github.com/dghubble/go-twitter/twitter)
* [oAuth1](https://pkg.go.dev/github.com/dghubble/oauth1)
* [Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
