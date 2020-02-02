package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"

	"github.com/pkg/errors"
)

func acquireFile(specFilename string, xmlSpecURL string) error {
	if fi, err := os.Stat(specFilename); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if fi.IsDir() {
			return errors.New("spec file is a directory")
		}
		return nil
	}

	resp, err := http.Get(xmlSpecURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("could not download spec %q from %q: %v",
			specFilename, xmlSpecURL, resp.Status)
	}

	tmpFilename := specFilename + ".download"
	w, err := os.Create(tmpFilename)
	if err != nil {
		return errors.Wrap(err, "creating temp file")
	}
	defer w.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return errors.Wrap(err, "downloading")
	}

	return os.Rename(tmpFilename, specFilename)
}

func globFiles(pattern string, archive *zip.ReadCloser) []*zip.File {
	var files []*zip.File
	for _, f := range archive.File {
		if matched, err := path.Match(pattern, f.Name); err != nil {
			// This shouldn't happen - all patterns are hard-coded, errors in them
			// are a programming error.
			panic(err)
		} else if matched {
			files = append(files, f)
		}
	}
	return files
}

func unmarshalXmlFile(file *zip.File, data interface{}) error {
	r, err := file.Open()
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(r)
	defer r.Close()
	return decoder.Decode(data)
}

var scpdFilenameRe = regexp.MustCompile(
	`.*/([a-zA-Z0-9]+)([0-9]+)\.xml`)

func urnPartsFromSCPDFilename(filename string) (*URNParts, error) {
	parts := scpdFilenameRe.FindStringSubmatch(filename)
	if len(parts) != 3 {
		return nil, fmt.Errorf("SCPD filename %q does not have expected number of parts", filename)
	}
	name, version := parts[1], parts[2]
	return &URNParts{
		Name:    name,
		Version: version,
	}, nil
}
