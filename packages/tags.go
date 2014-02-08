package packages

import (
	"code.google.com/p/goauth2/oauth"
	"encoding/json"
	"fmt"
	"github.com/rif/cache2go"
	"log"
	"os"
	"time"
)

var (
	HTTPClient = (&oauth.Transport{
		Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	}).Client()
)

type Tag struct {
	Name   string
	Commit Commit
}

type Commit struct {
	Sha string
	Url string
}

func GetTag(repo string, tag string) (Tag, error) {
	var t Tag
	var tags []Tag

	// check cache
	cache := cache2go.Cache("tags")
	item, err := cache.Value(repo)
	if err == nil {
		log.Println("Tag Cache Hit")
		tags = item.Data().([]Tag)
	} else {
		log.Println("Tag Cache miss")
		res, err := HTTPClient.Get("https://api.github.com/repos/" + repo + "/tags")
		if err != nil {
			return t, fmt.Errorf("Error contacting github")
		}

		log.Println("Github rate limit remaining", res.Header.Get("X-RateLimit-Remaining"))

		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&tags)
		if err != nil {
			return t, fmt.Errorf("Error Decoding JSON")
		}

		// cache the tags here
		cache.Cache(repo, 1*time.Minute, tags)
	}

	for _, val := range tags {
		if val.Name == tag {
			return val, nil
		}
	}

	return t, fmt.Errorf("Tag not found")
}
