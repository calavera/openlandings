package themes

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
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
