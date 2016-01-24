package themes

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

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
