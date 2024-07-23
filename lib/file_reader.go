package lib

import (
	"io"
	"os"
)

type FileReader struct{}

const PAGES_FOLDER = "./pages/"

func (fsr FileReader) Read(slug string) (string, error) {
	f, err := os.Open(PAGES_FOLDER + slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
