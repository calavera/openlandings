package themes

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/calavera/openlandings/github"
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

func Pack(repository *github.Repository, themePath, fullURL, landing string) (string, error) {
	site := Site{
		Title:       *repository.FullName,
		Description: *repository.Description,
		Content:     landing,
		BaseURL:     fullURL,
		Repo: Repo{
			Name:        *repository.FullName,
			Description: *repository.Description,
			URL:         *repository.HTMLURL,
		},
		Owner: Owner{
			Name:      *repository.Owner.Login,
			URL:       *repository.Owner.HTMLURL,
			AvatarURL: *repository.Owner.AvatarURL,
		},
	}

	idx, err := generateIndex(site, themePath)
	if err != nil {
		return "", err
	}

	zipFile, err := createZip(themePath, idx.Bytes())
	if err != nil {
		return "", err
	}

	return zipFile, nil
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

func generateIndex(site Site, themePath string) (*bytes.Buffer, error) {
	index := filepath.Join(themePath, "index.tpl")
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

func createZip(themePath string, idx []byte) (string, error) {
	deployDir, err := ioutil.TempDir("", "openlandings-zip-")
	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(filepath.Join(deployDir, "index.html"), idx, 0644); err != nil {
		return "", err
	}

	if err := writeStaticFiles(deployDir, themePath, "css"); err != nil {
		return "", err
	}

	if err := writeStaticFiles(deployDir, themePath, "fonts"); err != nil {
		return "", err
	}

	if err := writeStaticFiles(deployDir, themePath, "images"); err != nil {
		return "", err
	}

	if err := writeStaticFiles(deployDir, themePath, "js"); err != nil {
		return "", err
	}

	return deployDir, nil
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func writeStaticFiles(deployPath, sourcePath, name string) error {
	source := filepath.Join(sourcePath, name)
	dest := filepath.Join(deployPath, name)

	return copyFiles(source, dest)
}

func copyFiles(source, dest string) error {
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)
	for _, entry := range entries {

		sfp := filepath.Join(source, entry.Name())
		dfp := filepath.Join(dest, entry.Name())
		if entry.IsDir() {
			if err := copyFiles(sfp, dfp); err != nil {
				return err
			}
		} else {
			// perform copy
			if err := copyFile(sfp, dfp); err != nil {
				return err
			}
		}

	}
	return nil
}

func copyFile(source string, dest string) error {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
			return err
		}

	}

	return nil
}
