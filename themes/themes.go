package themes

import (
	"archive/zip"
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
	zipFile, err := ioutil.TempFile("", "openlandings-zip-")
	if err != nil {
		return "", err
	}

	w := zip.NewWriter(zipFile)
	f, err := w.Create("index.html")
	if err != nil {
		return "", err
	}
	if _, err := f.Write(idx); err != nil {
		return "", err
	}

	if err := writeStaticFiles(w, themePath, "css"); err != nil {
		return "", err
	}

	if err := writeStaticFiles(w, themePath, "fonts"); err != nil {
		return "", err
	}

	if err := writeStaticFiles(w, themePath, "images"); err != nil {
		return "", err
	}

	if err := writeStaticFiles(w, themePath, "js"); err != nil {
		return "", err
	}

	if err := w.Close(); err != nil {
		return "", err
	}

	return zipFile.Name(), nil
}

func writeStaticFiles(w *zip.Writer, themePath, dir string) error {
	base := filepath.Join(themePath, dir)
	_, err := os.Stat(base)
	if err == nil {
		if err := filepath.Walk(base, addFileToZip(base+"/", w)); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(base string, w *zip.Writer) filepath.WalkFunc {
	return func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			// do not add directories
			return nil
		}
		name := strings.TrimPrefix(base, path)
		f, err := w.Create(name)
		if err != nil {
			return err
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if _, err := f.Write(b); err != nil {
			return err
		}
		return nil
	}
}
