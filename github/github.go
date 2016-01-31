package github

import (
	"fmt"
	"io/ioutil"
	"strings"

	githubapi "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func newClient(accessToken string) *githubapi.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	ctx := newTransportContext()
	tc := oauth2.NewClient(ctx, ts)

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
	if err := checkMembership(client, owner); err != nil {
		return nil, err
	}
	orgOwner, _, err := client.Organizations.Get(owner)
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

func GetRepository(accessToken, login, nwo string) (*Repository, error) {
	parts := strings.SplitN(nwo, "/", 2)
	if len(parts) == 1 {
		return nil, fmt.Errorf("unable to recognize repository %s", nwo)
	}

	client := newClient(accessToken)
	if login != parts[0] {
		if err := checkMembership(client, parts[0]); err != nil {
			return nil, err
		}
	}

	repo, _, err := client.Repositories.Get(parts[0], parts[1])
	if err != nil {
		return nil, err
	}

	repository := &Repository{
		Repository: repo,
	}

	content, err := client.Repositories.DownloadContents(parts[0], parts[1], "README.md", nil)
	if err != nil {
		// ignore the content
		return repository, nil
	}
	defer content.Close()
	b, err := ioutil.ReadAll(content)
	if err != nil {
		// ignore the content
		return repository, nil
	}
	repository.Readme = string(b)

	return repository, nil
}

func checkMembership(client *githubapi.Client, owner string) error {
	// Ensure the user belongs to the org sent, or if it's even an org.
	_, _, err := client.Organizations.GetOrgMembership("", owner)
	if err != nil {
		return fmt.Errorf("the current user is not a member of the org %s: %v", owner, err)
	}

	return nil
}

type Repositories struct {
	Owner    *githubapi.Organization
	List     []githubapi.Repository
	NextPage int
	PrevPage int
}

type Repository struct {
	*githubapi.Repository
	Readme string
}
