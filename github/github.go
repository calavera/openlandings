package github

import "golang.org/x/oauth2"
import githubapi "github.com/google/go-github/github"

func newClient(accessToken string) *githubapi.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return githubapi.NewClient(tc)
}

func GetPrimaryEmail(accessToken string) string {
	client := newClient(accessToken)
	emails, _, err := client.Users.ListEmails(nil)
	var email string
	if err == nil && len(emails) > 0 {
		for _, e := range emails {
			if *e.Primary && *e.Verified {
				email = *e.Email
			}
		}
	}
	return email
}

func GetCurrentUser(accessToken string) (*githubapi.User, error) {
	client := newClient(accessToken)
	user, _, err := client.Users.Get("")
	return user, err
}

func ListOrganizations(accessToken string) ([]*githubapi.Organization, error) {
	client := newClient(accessToken)
	orgs, _, err := client.Organizations.List("", nil)

	var fullOrgs []*githubapi.Organization
	for _, org := range orgs {
		o, _, err := client.Organizations.Get(*org.Login)
		if err != nil {
			return nil, err
		}
		fullOrgs = append(fullOrgs, o)
	}

	return fullOrgs, err
}
