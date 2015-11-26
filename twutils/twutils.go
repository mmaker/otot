package twutils

import (
	"bufio"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func init() {
	twitterKey := os.Getenv("TWITTER_KEY")
	twitterSecret := os.Getenv("TWITTER_SECRET")
	if len(twitterKey) == 0 || len(twitterSecret) == 0 {
		panic(
			"Cannot get TWITTER_KEY and TWITTER_SECRET env variables. " +
			"Did you forget to set them?")
	}
	anaconda.SetConsumerKey(twitterKey)
	anaconda.SetConsumerSecret(twitterSecret)
}


func GetApi(credentials string) (api *anaconda.TwitterApi, err error) {
	f, err := os.Open(credentials)
	reader := bufio.NewReader(f)

	token, _, err := reader.ReadLine()
	secret, _, err := reader.ReadLine()

	api = anaconda.NewTwitterApi(string(token), string(secret))
	return
}
