package areena

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

type fixture struct {
	programPath string
	playoutPath string
}

func newFixture(id string) fixture {
	return fixture{
		programPath: fmt.Sprintf("./fixtures/%s.json", id),
		playoutPath: fmt.Sprintf("./fixtures/playout-%s.json", id),
	}
}

func readFile(relPath string) (contents []byte) {
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		panic(err)
	}

	contents, err = ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	return
}

// TODO use gobundle
func createMockAPIServer(fix fixture) *httptest.Server {
	programFixture := readFile(fix.programPath)
	playoutFixture := readFile(fix.playoutPath)
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/media/playouts.json" {
			fmt.Fprintln(w, string(playoutFixture))
			return
		}
		fmt.Fprintln(w, string(programFixture))
	}))
}
