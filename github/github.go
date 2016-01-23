package github

import (
	"fmt"

	"golang.org/x/oauth2"
)
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
	if err != nil {
		return nil, err
	}

	var fullOrgs []*githubapi.Organization
	for _, org := range orgs {
		o, _, err := client.Organizations.Get(*org.Login)
		if err != nil {
			return nil, err
		}
		fullOrgs = append(fullOrgs, o)
	}

	return fullOrgs, nil
}

func ListUserRepositories(accessToken string, opts *githubapi.RepositoryListOptions) (*Repositories, error) {
	client := newClient(accessToken)
	repos, resp, err := client.Repositories.List("", opts)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		List:     repos,
		NextPage: resp.NextPage,
		PrevPage: resp.PrevPage,
	}, nil
}

func ListOrgRepositories(accessToken, owner string, opts *githubapi.RepositoryListByOrgOptions) (*Repositories, error) {
	client := newClient(accessToken)

	var orgOwner *githubapi.Organization
	// Ensure the user belongs to the org sent, or if it's even an org.
	orgs, _, err := client.Organizations.List("", nil)
	if err != nil {
		return nil, err
	}
	var member bool
	for _, org := range orgs {
		if owner == *org.Login {
			member = true
		}
	}
	if !member {
		return nil, fmt.Errorf("the current user is not a member of the org %s", owner)
	}
	orgOwner, _, err = client.Organizations.Get(owner)
	if err != nil {
		return nil, err
	}

	repos, resp, err := client.Repositories.ListByOrg(owner, opts)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Owner:    orgOwner,
		List:     repos,
		NextPage: resp.NextPage,
		PrevPage: resp.PrevPage,
	}, nil
}

type Repositories struct {
	Owner    *githubapi.Organization
	List     []githubapi.Repository
	NextPage int
	PrevPage int
}
