package twitter

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type User struct {
	API *anaconda.TwitterApi
}

func InitTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("T_CK"))
	anaconda.SetConsumerSecret(os.Getenv("T_CS"))
	return anaconda.NewTwitterApi(os.Getenv("T_AT"), os.Getenv("T_ATS"))
}
