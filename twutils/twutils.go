package twutils

import (
	"bufio"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func init() {
	anaconda.SetConsumerKey("7kkQhns9JS0dY3cVpHHV5YqGv")
	anaconda.SetConsumerSecret("GWAJvLLmBaFSPWbgF6rohWKt93pgoytfNGSkuSMx1jVaZIw0e6")
}


func GetApi(credentials string) (api *anaconda.TwitterApi, err error) {
	f, err := os.Open(credentials)
	reader := bufio.NewReader(f)

	token, _, err := reader.ReadLine()
	secret, _, err := reader.ReadLine()

	api = anaconda.NewTwitterApi(string(token), string(secret))
	return
}
