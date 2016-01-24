package themes

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Owner struct {
	URL       string
	Name      string
	AvatarURL string
}

type Repo struct {
	Name        string
	Description string
	URL         string
}

type Site struct {
	Title       string
	Description string
	Content     string
	Analytics   string
	BaseURL     string
	Repo        Repo
	Owner       Owner
}

type Theme struct {
	Path string
}

func (t Theme) JSONFile() (io.Reader, error) {
	return os.Open(filepath.Join(t.Path, "info.json"))
}

func AllFromDisk() ([]Theme, error) {
	themesDir := os.Getenv("THEMES_DIR")
	if themesDir == "" {
		return nil, fmt.Errorf("THEMES_DIR empty, unable to find the templates")
	}

	var themes []Theme
	filepath.Walk(themesDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "/info.json") {
			themes = append(themes, Theme{filepath.Dir(path)})
		}
		return nil
	})

	return themes, nil
}

func Generate(site *Site, theme string) (io.Reader, error) {
	themesDir := os.Getenv("THEMES_DIR")
	if themesDir == "" {
		return nil, fmt.Errorf("THEMES_DIR empty, unable to find the templates")
	}

	index := filepath.Join(themesDir, theme, "index.tpl")
	tmpl, err := template.ParseFiles(index)
	if err != nil {
		return nil, fmt.Errorf("unable to read index template at %s: %v", index, err)
	}

	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, map[string]interface{}{"site": site}); err != nil {
		return nil, err
	}

	return buffer, nil
}
