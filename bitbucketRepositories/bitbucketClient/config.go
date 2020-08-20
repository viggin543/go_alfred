package bitbucketClient

import (
	"net/url"
	"os"
)

var user, pass string

func init() {
	user = url.QueryEscape(os.Getenv("BITBUCKET_USER"))
	pass = url.QueryEscape(os.Getenv("BITBUCKET_PASS"))
}