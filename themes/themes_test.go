package themes

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var testThemes = []string{
	"uno",
	"left",
	"base16",
}

func TestAllFromDisk(t *testing.T) {
	tmp, _ := ioutil.TempDir("", "ol-")
	uno := filepath.Join(tmp, "uno")
	os.Mkdir(uno, 0755)
	defer os.RemoveAll(tmp)

	info := filepath.Join(uno, "info.json")
	ioutil.WriteFile(info, []byte(`{"name":"uno", "description":"uno", "home":"uno"}`), 644)

	os.Setenv("THEMES_DIR", tmp)
	defer os.Setenv("THEMES_DIR", "")

	th, err := AllFromDisk()
	if err != nil {
		t.Fatal(err)
	}
	if len(th) != 1 {
		t.Fatalf("expected 1, got %v", len(th))
	}
}

func TestGenerateIndex(t *testing.T) {
	base, _ := os.Getwd()
	themesDir := filepath.Join(base, "../static/themes")

	site := Site{
		Title:       "test title",
		Description: "test description",
		Content:     template.HTML("test content"),
		BaseURL:     "http://localhost:1234",
		Repo: Repo{
			Name:        "test repo",
			Description: "test description",
			Login:       "test/repo",
			URL:         "https://github.com/test/repo",
		},
		Owner: Owner{
			Name:      "test",
			URL:       "https://github.com/test",
			AvatarURL: "https://avatars.github.com/test",
		},
	}

	for _, theme := range testThemes {
		themePath := filepath.Join(themesDir, theme)
		_, err := generateIndex(site, themePath)
		if err != nil {
			t.Fatal(err)
		}
	}
}
