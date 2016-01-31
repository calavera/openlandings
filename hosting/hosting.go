package hosting

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/calavera/openlandings/models"
	"github.com/netlify/netlify-go"
)

func Publish(user models.User, zipFilePath, siteName, domain string) error {
	client, err := newUserClient(user)
	if err != nil {
		return err
	}

	site, _, err := client.Sites.Create(&netlify.SiteAttributes{
		Name:              siteName,
		CustomDomain:      domain,
		NotificationEmail: user.Email,
	})
	if err != nil {
		return err
	}

	// Enable TLS
	if domain != "" {
		_, err := site.ProvisionCert(nil)
		if err != nil {
			return err
		}
		site.ForceSSL = true
		_, err = site.Update()
		if err != nil {
			return err
		}
	}

	// Deploy a directory
	deploy, _, err := site.Deploys.Create(zipFilePath)
	if err != nil {
		return err
	}

	// Wait for the deploy to process
	return deploy.WaitForReady(0)
}

func Delete(site *models.Site, user models.User) error {
	client, err := newUserClient(user)
	if err != nil {
		return err
	}

	d, err := url.Parse(site.Domain)
	if err != nil {
		return err
	}

	hosted, _, err := client.Sites.Get(d.Host)
	if err != nil {
		return err
	}

	_, err = hosted.Destroy()
	return err
}

func newUserClient(user models.User) (*netlify.Client, error) {
	accessToken, err := getUserAccessToken(user.Email, user.UUID)
	if err != nil {
		return nil, err
	}

	return netlify.NewClient(&netlify.Config{
		UserAgent:   "openlandings.com",
		AccessToken: accessToken,
	}), nil
}

// POST https://api.netlify.com/api/v1/access_tokens
// {"user": {"email": “name@mail.com", "uid": "12345"}}
// Returns:
// {"id":"54321","access_token":"some-token","user_id":"54321","created_at":"2015-11-06T02:48:15Z","email":"name@email.com"}
func getUserAccessToken(email, uuid string) (string, error) {
	resellerAccessToken := os.Getenv("RESELLER_KEY")
	if resellerAccessToken == "" {
		return "", fmt.Errorf("unable to publish site without reseller information")
	}

	client := netlify.NewClient(&netlify.Config{
		UserAgent:   "openlandings.com",
		AccessToken: resellerAccessToken,
	})

	netlifyUser := &netlifyUser{
		Email: email,
		UUID:  uuid,
	}
	opts := &netlify.RequestOptions{
		JsonBody: map[string]interface{}{"user": netlifyUser},
	}

	resp, err := client.Request("POST", "/access_tokens", opts, netlifyUser)
	if err != nil {
		b, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error creating user %s: %v\n%s", email, err, b)
	}

	if netlifyUser.AccessToken == "" {
		return "", fmt.Errorf("unable to authenticate user with email %s", email)
	}
	return netlifyUser.AccessToken, nil
}

type netlifyUser struct {
	Email       string `json:"email"`
	UUID        string `json:"uid"`
	AccessToken string `json:"access_token,omitempty"`
}
