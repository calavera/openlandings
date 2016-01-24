package hosting

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/calavera/openlandings/models"
	"github.com/netlify/netlify-go"
)

func Publish(user models.User, zipFilePath, siteName, domain string) error {
	accessToken, err := getUserAccessToken(user)
	if err != nil {
		return err
	}

	client := netlify.NewClient(&netlify.Config{
		UserAgent:   "openlandings.com",
		AccessToken: accessToken,
	})

	site, _, err := client.Sites.Create(&netlify.SiteAttributes{
		Name:              siteName,
		CustomDomain:      domain,
		Password:          user.SitesPassword,
		NotificationEmail: user.Email,
	})
	if err != nil {
		return err
	}

	// Deploy a directory
	deploy, _, err := site.Deploys.Create(zipFilePath)
	if err != nil {
		return err
	}

	// Wait for the deploy to process
	return deploy.WaitForReady(0)
}

// POST https://api.netlify.com/api/v1/access_tokens
// {"user": {"email": “name@mail.com", "uid": "12345"}}
// Returns:
// {"id":"54321","access_token":"some-token","user_id":"54321","created_at":"2015-11-06T02:48:15Z","email":"name@email.com"}
func getUserAccessToken(user models.User) (string, error) {
	resellerAccessToken := os.Getenv("RESELLER_KEY")
	if resellerAccessToken == "" {
		return "", fmt.Errorf("unable to publish site without reseller information")
	}

	client := netlify.NewClient(&netlify.Config{
		UserAgent:   "openlandings.com",
		AccessToken: resellerAccessToken,
	})

	netlifyUser := &netlifyUser{
		User: user.Email,
		UUID: user.UUID,
	}
	opts := &netlify.RequestOptions{
		JsonBody: netlifyUser,
	}

	resp, err := client.Request("POST", "/access_token", opts, netlifyUser)
	if err != nil {
		b, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error creating user %s: %v\n%s", user.Email, err, b)
	}

	if netlifyUser.AccessToken == "" {
		return "", fmt.Errorf("unable to authenticate user with email %s", user.Email)
	}
	return netlifyUser.AccessToken, nil
}

type netlifyUser struct {
	User        string `json:"user"`
	UUID        string `json:"uid"`
	AccessToken string `json:"access_token,omitempty"`
}
